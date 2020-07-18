package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	bs_media "picasso/business/service/media"
)

type UploadedMediaSliceHashes struct {
	ghost.ApiTemplate
}

func (this *UploadedMediaSliceHashes) Resource() string{
	return "upload.uploaded_media_slice_hashes"
}

type uploadedMediaSliceHashesGetParams struct {
	CompleteFilename string `form:"complete_filename"`
	CompleteHash string `form:"complete_hash"`
	TotalSliceCount int `form:"total_slice_count"`
	SLiceHashes string `form:"slice_hashes"`
	SliceHash2index string `form:"slice_hash2index"`
}
// 支持秒传
func (this *UploadedMediaSliceHashes) Get() ghost.Response{
	ctx := this.GetCtx()
	var params uploadedMediaSliceHashesGetParams
	this.Bind(&params)
	var hashList []string
	ghost_util.Decode(params.SLiceHashes, &hashList)
	var sliceHash2index map[string]int
	ghost_util.Decode(params.SliceHash2index, &sliceHash2index)
	return ghost.NewJsonResponse(
		bs_media.NewMediaService(ctx).CheckSliceExistenceByHashes(
			params.CompleteFilename, params.CompleteHash, params.TotalSliceCount,
			hashList, sliceHash2index),
	)
}

func init(){
	ghost.RegisterApi(&UploadedMediaSliceHashes{})
}
