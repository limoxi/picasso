package upload

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	bs_media "picasso/business/service/media"

	//bs_space "picasso/business/service/space"
)

type UploadedSlicedMedia struct {
	ghost.ApiTemplate
}

func (this *UploadedSlicedMedia) Resource() string{
	return "upload.uploaded_sliced_media"
}

type uploadedSlicedMediaPutParams struct {
	SpaceId int `form:"space_id"`
	Filename string `form:"filename"`
	CompleteHash string `form:"complete_hash"`
	SliceHash string `form:"slice_hash"`
	SliceIndex int `form:"slice_index"`
	Slice *multipart.FileHeader `form:"slice"`
	TotalSliceCount int `form:"total_slice_count"`
}

func (this *UploadedSlicedMedia) Put() ghost.Response{
	ctx := this.GetCtx()
	var params uploadedSlicedMediaPutParams
	this.Bind(&params)

	bs_media.NewUploader(ctx).UploadSlicedMedia(
		params.SpaceId,
		params.Filename,
		params.Slice,
		params.CompleteHash,
		params.SliceHash,
		params.SliceIndex,
		params.TotalSliceCount,
	)
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&UploadedSlicedMedia{})
}
