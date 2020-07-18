package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	bs_media "picasso/business/service/media"
)

type UploadedMediaHashes struct {
	ghost.ApiTemplate
}

func (this *UploadedMediaHashes) Resource() string{
	return "upload.uploaded_media_hashes"
}

type uploadedMediaHashesGetParams struct {
	Hashes string `form:"hashes"`
}
// 支持秒传，检查hashcode是否已存在
func (this *UploadedMediaHashes) Get() ghost.Response{
	ctx := this.GetCtx()
	var params uploadedMediaHashesGetParams
	this.Bind(&params)
	var hashList []string
	ghost_util.Decode(params.Hashes, &hashList)
	return ghost.NewJsonResponse(
		bs_media.NewMediaService(ctx).CheckExistenceByHashes(hashList),
	)
}

func init(){
	ghost.RegisterApi(&UploadedMediaHashes{})
}
