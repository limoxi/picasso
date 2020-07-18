package upload

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	bs_media "picasso/business/service/media"
)

type UploadedImages struct {
	ghost.ApiTemplate
}

func (this *UploadedImages) Resource() string{
	return "upload.uploaded_images"
}

type uploadImagesPutParams struct {
	SpaceId int `form:"space_id"`
	Files []*multipart.FileHeader `form:"files[]"`
}
// 批量上传图片
func (this *UploadedImages) Put() ghost.Response{
	ctx := this.GetCtx()
	var params uploadImagesPutParams
	this.Bind(&params)

	bs_media.NewUploader(ctx).UploadImages(params.SpaceId, params.Files)
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&UploadedImages{})
}
