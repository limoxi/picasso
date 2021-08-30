package upload

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	bs_media "picasso/business/app/media"
	db_media "picasso/db/media"

)

type UploadedSlicedMedia struct {
	ghost.ApiTemplate

	PutParams *struct{
		MediaType string `form:"media_type"`
		SpaceId int `form:"space_id"`
		Filename string `form:"filename"`
		CompleteHash string `form:"complete_hash"`
		SliceHash string `form:"slice_hash"`
		SliceIndex int `form:"slice_index"`
		Slice *multipart.FileHeader `form:"slice"`
		TotalSliceCount int `form:"total_slice_count"`
	}
}

func (this *UploadedSlicedMedia) Resource() string{
	return "upload.uploaded_sliced_media"
}

func (this *UploadedSlicedMedia) Put() ghost.Response{
	ctx := this.GetCtx()
	params := this.PutParams

	bs_media.NewUploader(ctx).UploadSlicedMedia(&bs_media.SliceUploadParams{
		MediaType:       db_media.MEDIA_TEXT2TYPE[params.MediaType],
		SpaceId:         params.SpaceId,
		Filename:        params.Filename,
		FileHeader:      params.Slice,
		CompleteHash:    params.CompleteHash,
		SliceHash:       params.SliceHash,
		SliceIndex:      params.SliceIndex,
		TotalSliceCount: params.TotalSliceCount,
	})
	return ghost.NewJsonResponse(nil)
}

func init(){
	ghost.RegisterApi(&UploadedSlicedMedia{})
}
