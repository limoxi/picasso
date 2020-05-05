package media

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"github.com/limoxi/ghost"
	"io"
	"mime/multipart"
	"os"
	"path"
)

var STORAGE_ROOT_PATH string
var IMAGE_STORAGE_PATH string
var VIDEO_STORAGE_PATH string
var THUMBNAIL_STORAGE_PATH string

type Uploader interface {
	Upload(spaceId int, fileHeader *multipart.FileHeader, fileHash string)
}

// CheckFileHash hash匹配
func CheckFileHash(fh *multipart.FileHeader, fileHash string) bool{
	hash := md5.New()
	var content []byte
	f, _ := fh.Open()
	_, err := f.Read(content)
	if err != nil{
		ghost.Error(err)
		panic(ghost.NewSystemError("读取文件内容失败"))
	}
	hash.Write(content)
	s := make([]byte, hex.EncodedLen(hash.Size()))
	hex.Encode(s, hash.Sum(nil))
	ghost.Info(string(bytes.ToLower(s)))
	return fileHash == string(bytes.ToLower(s))
}

func SaveFile(f multipart.File, dst string) error{
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, f)
	return err
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