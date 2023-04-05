package file

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
	"picasso/common/util"
	db_file "picasso/db/file"
	"strings"
)

type UploadParams struct {
	User        *bm_account.User
	Path        string
	FileHeaders []*multipart.FileHeader
	Hashes      []string
}

type SliceUploadParams struct {
	Path            string
	Filename        string
	FileHeader      *multipart.FileHeader
	CompleteHash    string
	SliceHash       string
	SliceIndex      int
	TotalSliceCount int
}

type UploadResult map[string]bool

type fileInfo struct {
	Hash      string
	MimeType  string
	Ext       string
	MediaType int
}

type Uploader struct {
	ghost.DomainService
}

func (this *Uploader) GetHash(fh *multipart.FileHeader) (string, error) {
	hashCode := ""
	f, err := fh.Open()
	if err != nil {
		ghost.Error(err)
		return hashCode, ghost.NewSystemError("打开文件失败")
	}
	defer f.Close()

	hash := md5.New()
	content := make([]byte, fh.Size)
	_, err = f.Read(content)
	if err != nil {
		ghost.Error(err)
		return hashCode, ghost.NewSystemError("读取文件内容失败")
	}
	hash.Write(content)
	s := make([]byte, hex.EncodedLen(hash.Size()))
	hex.Encode(s, hash.Sum(nil))
	return string(bytes.ToLower(s)), nil
}

func (this *Uploader) getGroupForUser(groupId int, user *bm_account.User) int {
	if groupId == 0 {
		group := bm_file.NewGroupRepository(this.GetCtx()).GetDefaultForUser(user)
		return group.Id
	}
	groupUser := bm_file.NewGroupUserRepository(this.GetCtx()).GetById(groupId, user.Id)
	if groupUser == nil || !groupUser.IsManager {
		panic(ghost.NewBusinessError("当前用户无权限上传文件"))
	}
	return groupId
}

func (this *Uploader) saveFile(fh *multipart.FileHeader, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	f, err := fh.Open()
	if err != nil {
		ghost.Error(err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(out, f)
	return err
}

func (this *Uploader) checkPath(path string) {
	c := strings.ToUpper(strings.Split(path, ".")[0])
	if db_file.CATEGORY2CN[c] == "" {
		panic(ghost.NewBusinessError("不支持的类别"))
	}
}

func (this *Uploader) UploadFiles(params *UploadParams) UploadResult {
	uploadResult := make(UploadResult)
	if len(params.FileHeaders) == 0 {
		return uploadResult
	}

	this.checkPath(params.Path)

	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx)
	hashes := make([]string, 0)
	hash2fh := make(map[string]*multipart.FileHeader)
	for index, fh := range params.FileHeaders {
		eHash := params.Hashes[index]
		uploadResult[eHash] = false

		if fh.Size == 0 {
			ghost.Error("file size is 0")
			continue
		}
		hash, err := this.GetHash(fh)
		if err != nil {
			ghost.Error("calc hash code failed")
			continue
		}

		if hash != eHash {
			panic(ghost.NewBusinessError("invalid_hash",
				fmt.Sprintf("文件hash不一致:%s-%s!=%s", fh.Filename, hash, eHash)))
		}
		hashes = append(hashes, hash)
		hash2fh[hash] = fh
	}
	var dbModels []*db_file.File
	result := db.Model(&db_file.File{}).Where(ghost.Map{
		"hash__in": hashes,
	}).Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}
	existedHashes := make([]string, 0)
	for _, dbModel := range dbModels {
		existedHashes = append(existedHashes, dbModel.Hash)
	}

	lister := ghost_util.NewListerFromStrings(existedHashes)
	for _, hash := range hashes {
		if lister.Has(hash) {
			uploadResult[hash] = true
			continue
		}
		fh := hash2fh[hash]

		storagePath := path.Join(
			STORAGE_ROOT_PATH,
			strings.ReplaceAll(params.Path, ".", string(os.PathSeparator)),
			fh.Filename)
		err := this.saveFile(fh, storagePath)
		if err != nil {
			ghost.Error(err)
			continue
		}
		result := db.Create(&db_file.File{
			UserId:      params.User.Id,
			Type:        db_file.FILE_TYPE_FILE,
			Path:        params.Path,
			Hash:        hash,
			Name:        fh.Filename,
			Size:        fh.Size,
			Status:      db_file.FILE_STATUS_SAVED,
			CreatedTime: ghost_util.DEFAULT_TIME,
		})
		if err := result.Error; err != nil {
			ghost.Error(err)
			continue
		}
		uploadResult[hash] = true
	}

	return uploadResult
}

func (this *Uploader) slicedMediaIsComplete(dirPath, pureFilename, hash string, totalSliceCount int) bool {
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	if len(fs) <= totalSliceCount {
		return false
	}
	// 生成群序列文件名
	allFilenames := make([]string, 0, totalSliceCount)
	for _, f := range fs {
		allFilenames = append(allFilenames, f.Name())
	}
	lister := ghost_util.NewListerFromStrings(allFilenames)
	for i := 0; i < totalSliceCount; i++ {
		if !lister.Has(fmt.Sprintf("%s_%s_%s.slice",
			pureFilename, util.WrappedInt(totalSliceCount), util.WrappedInt(i))) {
			return false
		}
	}
	return true
}

// UploadSlicedFile
// 文件格式 blockIndex.blockCount.filename.hash.sliced
func (this *Uploader) UploadSlicedFile(params *SliceUploadParams) {
	this.checkPath(params.Path)
	h, err := this.GetHash(params.FileHeader)
	if err != nil {
		panic(err)
	}
	if params.SliceHash != h {
		panic(ghost.NewBusinessError("invalid_hash", fmt.Sprintf("文件hash不一致:%s-%s", params.SliceHash, h)))
	}

	pureFileName := strings.Split(params.Filename, ".")[0]
	dirPath := fmt.Sprintf("tmp_%s_%s", pureFileName, params.CompleteHash)
	tmpDirPath := path.Join(SLICE_TMP_STORAGE_PATH, dirPath)
	err = os.Mkdir(tmpDirPath, os.ModeDir)
	if err != nil {
		ghost.Warn(err)
	}
	sliceFilename := fmt.Sprintf("%s_%s_%s.slice",
		pureFileName, util.WrappedInt(params.TotalSliceCount), util.WrappedInt(params.SliceIndex))
	storagePath := path.Join(tmpDirPath, sliceFilename)
	err = this.saveFile(params.FileHeader, storagePath)
	if err != nil {
		panic(err)
	}
	db := ghost.GetDBFromCtx(this.GetCtx())
	dbModel := &db_file.File{
		Path:        params.Path,
		Type:        db_file.FILE_TYPE_FILE,
		Hash:        params.CompleteHash,
		Name:        params.Filename,
		Status:      db_file.FILE_STATUS_SLICE_SAVED,
		CreatedTime: ghost_util.DEFAULT_TIME,
	}
	var count int64
	if err = db.Model(&db_file.File{}).Where(ghost.Map{
		"hash": params.CompleteHash,
	}).Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 0 {
		result := db.Create(dbModel)
		if err := result.Error; err != nil {
			panic(err)
		}
	}

	result := db.Create(&db_file.SlicedFile{
		FileHash:        params.CompleteHash,
		SliceHash:       params.SliceHash,
		SliceIndex:      params.SliceIndex,
		TotalSliceCount: params.TotalSliceCount,
		Path:            storagePath,
		Size:            params.FileHeader.Size,
		Status:          db_file.SLICED_FILE_STATUS_SAVED,
	})
	if err := result.Error; err != nil {
		panic(err)
	}
}

func NewUploader(ctx context.Context) *Uploader {
	inst := new(Uploader)
	inst.SetCtx(ctx)
	return inst
}
