package media

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	"mime/multipart"
	"os"
	"path"
	ghost_util "github.com/limoxi/ghost/utils"
	db_media "picasso/db/media"
)

type ImageUploadService struct {
	ghost.DomainObject
}

// Upload
// 1、存储文件到指定目录
// 2、根据文件名和大小生成hash
// 3、将文件信息存入数据库
// 4、交由cron协程，识别exif信息
func (this *ImageUploadService) Upload(spaceId int, fileHeader *multipart.FileHeader, fileHash string) {
	ctx := this.GetCtx()
	ginCtx := ctx.(*gin.Context)
	db := ghost.GetDBFromCtx(ctx)
	if !CheckFileHash(fileHeader, fileHash){
		panic(ghost.NewBusinessError("文件已存在"))
	}
	storagePath := path.Join(IMAGE_STORAGE_PATH, string(os.PathSeparator), fileHeader.Filename)
	result := db.Create(&db_media.Media{
		SpaceId: spaceId,
		Type: db_media.MEDIA_TYPE_IMAGE,
		Hash: fileHash,
		StoragePath: storagePath,
		Status: db_media.MEDIA_STATUS_SAVED,
		Size: fileHeader.Size,
		ShootTime: ghost_util.DEFAULT_TIME,
	})
	if err := result.Error; err != nil{
		ghost.Error(err)
		panic(err)
	}

	err := ginCtx.SaveUploadedFile(fileHeader, storagePath)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
}

func NewImageUploadService(ctx context.Context) *ImageUploadService{
	inst := new(ImageUploadService)
	inst.SetCtx(ctx)
	return inst
}