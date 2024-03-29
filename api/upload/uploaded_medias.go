package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"mime/multipart"
	bs_media "picasso/business/app/media"
	db_media "picasso/db/media"
)

type UploadedMedias struct {
	ghost.ApiTemplate

	PutParams *struct{
		SpaceId int `form:"space_id"`
		MediaType string `form:"media_type"`
		Files []*multipart.FileHeader `form:"files[]"`
		Filename2Hash string `form:"filename2hash"`
	}
}

func (this *UploadedMedias) Resource() string{
	return "upload.uploaded_medias"
}

// Put 批量上传
func (this *UploadedMedias) Put() ghost.Response{
	ctx := this.GetCtx()
	params := this.PutParams

	var filename2hash map[string]string
	ghost_util.Decode(params.Filename2Hash, &filename2hash)
	bs_media.NewUploader(ctx).UploadMedias(&bs_media.UploadParams{
		MediaType: db_media.MEDIA_TEXT2TYPE[params.MediaType],
		SpaceId: params.SpaceId,
		FileHeaders: params.Files,
		Filename2Hash: filename2hash,
	})
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&UploadedMedias{})
}
