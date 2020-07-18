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

var STORAGE_ROOT_PATH string
var IMAGE_STORAGE_PATH string
var VIDEO_STORAGE_PATH string
var THUMBNAIL_STORAGE_PATH string

type UploadParams struct {
	MediaType int
	SpaceId int
	Hash string
}

type SliceUploadParams struct {
	UploadParams
	BlobNum int
	TotalBlobNum int
	BlobSize int
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

func (this *Uploader) UploadImages(spaceId int, fhs []*multipart.FileHeader){
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx)
	hashes := make([]string, 0)
	hash2fh := make(map[string]*multipart.FileHeader)
	for _, fh := range fhs{
		if fh.Size == 0{
			continue
		}
		hash, err := this.GetHash(fh)
		if err != nil{
			ghost.Error("calc hash code failed")
			continue
		}
		hashes = append(hashes, hash)
		hash2fh[hash] = fh
	}
	var dbModels []*db_media.Media
	result := db.Model(&db_media.Media{}).Where(ghost.Map{
		"hash__in":  hashes,
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

		storagePath := path.Join(IMAGE_STORAGE_PATH, string(os.PathSeparator), fh.Filename)
		err := this.saveFile(fh, storagePath)
		if err != nil{
			ghost.Error(err)
			continue
		}
		result := db.Create(&db_media.Media{
			SpaceId: spaceId,
			Type: db_media.MEDIA_TYPE_IMAGE,
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
func (this *Uploader) UploadSlicedMedia(spaceId int, filename string, fh *multipart.FileHeader,
	completeHash, sliceHash string, sliceIndex, totalSliceCount int){

	h, err := this.GetHash(fh)
	if err != nil{
		panic(err)
	}
	if sliceHash != h{
		panic(ghost.NewBusinessError("invalid_hash", fmt.Sprintf("文件hash不一致:%s-%s", sliceHash, h)))
	}
	pureFileName := strings.Split(filename, ".")[0]
	tmpDirPath := path.Join(VIDEO_STORAGE_PATH, fmt.Sprintf("tmp_%s_%s", pureFileName, completeHash))
	err = os.Mkdir(tmpDirPath, os.ModeDir)
	if err != nil{
		ghost.Warn(err)
	}
	sliceFilename := fmt.Sprintf("%s_%s_%d_%d.slice",
		pureFileName, sliceHash, totalSliceCount, sliceIndex)
	storagePath := path.Join(tmpDirPath, sliceFilename)
	err = this.saveFile(fh, storagePath)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
	result := ghost.GetDBFromCtx(this.GetCtx()).Create(&db_media.Media{
		SpaceId: spaceId,
		Type: db_media.MEDIA_TYPE_VIDEO,
		Hash: completeHash,
		Status: db_media.MEDIA_STATUS_SLICE_SAVED,
		ShootTime: ghost_util.DEFAULT_TIME,
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

func init(){
	if ghost.OS == "windows"{
		STORAGE_ROOT_PATH = "E:\\picasso"
	}else{
		STORAGE_ROOT_PATH = "/picasso"
	}
	IMAGE_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, string(os.PathSeparator), "image")
	VIDEO_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, string(os.PathSeparator), "video")
	THUMBNAIL_STORAGE_PATH = path.Join(STORAGE_ROOT_PATH, string(os.PathSeparator), "thumbnail")
}