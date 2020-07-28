package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	bs_media "picasso/business/service/media"
	db_media "picasso/db/media"
)

type UploadedMediaSliceHashes struct {
	ghost.ApiTemplate
}

func (this *UploadedMediaSliceHashes) Resource() string{
	return "upload.uploaded_media_slice_hashes"
}

type uploadedMediaSliceHashesGetParams struct {
	MediaType string `form:"media_type"`
	CompleteHash string `form:"complete_hash"`
	SLiceHashes string `form:"slice_hashes"`
}
// 支持秒传
func (this *UploadedMediaSliceHashes) Get() ghost.Response{
	ctx := this.GetCtx()
	var params uploadedMediaSliceHashesGetParams
	this.Bind(&params)
	var hashList []string
	ghost_util.Decode(params.SLiceHashes, &hashList)
	return ghost.NewJsonResponse(
		bs_media.NewMediaService(ctx).CheckSliceExistenceByHashes(
			db_media.MEDIA_TEXT2TYPE[params.MediaType], params.CompleteHash, hashList),
	)
}

func init(){
	ghost.RegisterApi(&UploadedMediaSliceHashes{})
}
