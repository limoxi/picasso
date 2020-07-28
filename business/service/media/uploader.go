package media

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
	db_media "picasso/db/media"
	"strings"
)

const MAX_MEDIA_SIZE = 32 * 1024 * 1024 // 单个文件最大32m

var STORAGE_ROOT_PATH string
var IMAGE_STORAGE_PATH string
var VIDEO_STORAGE_PATH string
var THUMBNAIL_STORAGE_PATH string
var MEDIA_TYPE2STORAGE_PATH map[int]string

type UploadParams struct {
	MediaType int
	SpaceId int
	FileHeaders []*multipart.FileHeader
	Filename2Hash map[string]string
}

type SliceUploadParams struct {
	MediaType int
	SpaceId int
	Filename string
	FileHeader *multipart.FileHeader
	CompleteHash string
	SliceHash string
	SliceIndex int
	TotalSliceCount int
}

type Uploader struct {
	ghost.DomainService
}

func (this *Uploader) GetHash(fh *multipart.FileHeader) (string, error) {
	hashCode := ""
	f, err := fh.Open()
	if err != nil{
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

func (this *Uploader) saveFile(fh *multipart.FileHeader, path string) error{
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	f, err := fh.Open()
	if err != nil{
		ghost.Error(err)
		return err
	}
	defer f.Close()

	_, err = io.Copy(out, f)
	return err
}

func (this *Uploader) UploadMedias(params *UploadParams){
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx)
	hashes := make([]string, 0)
	hash2fh := make(map[string]*multipart.FileHeader)
	for _, fh := range params.FileHeaders{
		if fh.Size == 0{
			continue
		}
		hash, err := this.GetHash(fh)
		if err != nil{
			ghost.Error("calc hash code failed")
			continue
		}
		if hash != params.Filename2Hash[fh.Filename]{
			panic(ghost.NewBusinessError("invalid_hash",
				fmt.Sprintf("文件hash不一致:%s-%s", params.Filename2Hash[fh.Filename], hash)))
		}
		hashes = append(hashes, hash)
		hash2fh[hash] = fh
	}
	var dbModels []*db_media.Media
	result := db.Model(&db_media.Media{}).Where(ghost.Map{
		"hash__in":  hashes,
		"type": params.MediaType,
	}).Find(&dbModels)
	if err := result.Error; err != nil{
		ghost.Error(err)
		panic(err)
	}
	existedHashes := make([]string, 0)
	for _, dbModel := range dbModels{
		existedHashes = append(existedHashes, dbModel.Hash)
	}
	lister := ghost_util.NewListerFromStrings(existedHashes)
	for _, hash := range hashes{
		if lister.Has(hash){
			continue
		}
		fh := hash2fh[hash]

		storagePath := path.Join(MEDIA_TYPE2STORAGE_PATH[params.MediaType], string(os.PathSeparator), fh.Filename)
		err := this.saveFile(fh, storagePath)
		if err != nil{
			ghost.Error(err)
			continue
		}
		result := db.Create(&db_media.Media{
			SpaceId: params.SpaceId,
			Type: params.MediaType,
			Hash: hash,
			StoragePath: storagePath,
			Status: db_media.MEDIA_STATUS_SAVED,
			Size: fh.Size,
			ShootTime: ghost_util.DEFAULT_TIME,
		})
		if err := result.Error; err != nil{
			ghost.Error(err)
			panic(err)
		}
	}
}

func (this *Uploader) slicedMediaIsComplete(dirPath, pureFilename, hash string, totalSliceCount int) bool{
	fs, err := ioutil.ReadDir(dirPath)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
	if len(fs) <= totalSliceCount{
		return false
	}
	// 生成群序列文件名
	allFilenames := make([]string, 0, totalSliceCount)
	for _, f := range fs{
		allFilenames = append(allFilenames, f.Name())
	}
	lister := ghost_util.NewListerFromStrings(allFilenames)
	for i:=0; i<totalSliceCount; i++{
		if !lister.Has(fmt.Sprintf("%s_%s_%d_%d.slice",
			pureFilename, hash, totalSliceCount, i)){
			return false
		}
	}
	return true
}

// 文件格式 blockIndex_blockCount_filename_hash.sliced
func (this *Uploader) UploadSlicedMedia(params *SliceUploadParams){

	h, err := this.GetHash(params.FileHeader)
	if err != nil{
		panic(err)
	}
	if params.SliceHash != h{
		panic(ghost.NewBusinessError("invalid_hash", fmt.Sprintf("文件hash不一致:%s-%s", params.SliceHash, h)))
	}
	pureFileName := strings.Split(params.Filename, ".")[0]
	tmpDirPath := path.Join(MEDIA_TYPE2STORAGE_PATH[params.MediaType], fmt.Sprintf("tmp_%s_%s", pureFileName, params.CompleteHash))
	err = os.Mkdir(tmpDirPath, os.ModeDir)
	if err != nil{
		ghost.Warn(err)
	}
	sliceFilename := fmt.Sprintf("%s_%s_%d_%d.slice",
		pureFileName, params.SliceHash, params.TotalSliceCount, params.SliceIndex)
	storagePath := path.Join(tmpDirPath, sliceFilename)
	err = this.saveFile(params.FileHeader, storagePath)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
	db := ghost.GetDBFromCtx(this.GetCtx())
	dbModel := &db_media.Media{
		SpaceId: params.SpaceId,
		Type: params.MediaType,
		Hash: params.CompleteHash,
		Status: db_media.MEDIA_STATUS_SLICE_SAVED,
		ShootTime: ghost_util.DEFAULT_TIME,
	}
	result := db.Create(dbModel)
	if err := result.Error; err != nil{
		ghost.Error(err)
		panic(err)
	}

	result = db.Create(&db_media.SlicedMedia{
		MediaType:   params.MediaType,
		MediaHash:   params.CompleteHash,
		SliceHash:   params.SliceHash,
		SliceIndex:  params.SliceIndex,
		TotalSliceCount: params.TotalSliceCount,
		StoragePath: storagePath,
		Status:      db_media.SLICED_MEDIA_STATUS_SAVED,
		Size:        params.FileHeader.Size,
	})
	if err := result.Error; err != nil{
		ghost.Error(err)
		panic(err)
	}
}

func NewUploader(ctx context.Context) *Uploader {
	inst := new(Uploader)
	inst.SetCtx(ctx)
	return inst
}

func prepareDirs(){
	err := os.Mkdir(IMAGE_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(VIDEO_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)

	err = os.Mkdir(THUMBNAIL_STORAGE_PATH, os.ModeDir)
	ghost.Warn(err)
}

func init(){
	if ghost.OS == "windows"{
		STORAGE_ROOT_PATH = "E:\\picasso"
	}else{
		STORAGE_ROOT_PATH = "/picasso"
	}
	IMAGE_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, string(os.PathSeparator), "image")
	VIDEO_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, string(os.PathSeparator), "video")
	THUMBNAIL_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, string(os.PathSeparator), "thumbnail")

	MEDIA_TYPE2STORAGE_PATH = map[int]string{
		db_media.MEDIA_TYPE_IMAGE: IMAGE_STORAGE_PATH,
		db_media.MEDIA_TYPE_VIDEO: VIDEO_STORAGE_PATH,
	}

	prepareDirs()
}