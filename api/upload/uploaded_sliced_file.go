package upload

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	bs_media "picasso/business/app/file"
	db_file "picasso/db/file"
)

type UploadedSlicedFile struct {
	ghost.ApiTemplate

	PutParams *struct {
		FileType        string                `form:"file_type"`
		GroupId         int                   `form:"group_id"`
		Filename        string                `form:"filename"`
		CompleteHash    string                `form:"complete_hash"`
		SliceHash       string                `form:"slice_hash"`
		SliceIndex      int                   `form:"slice_index"`
		Slice           *multipart.FileHeader `form:"slice"`
		TotalSliceCount int                   `form:"total_slice_count"`
	}
}

func (this *UploadedSlicedFile) Resource() string {
	return "upload.uploaded_sliced_file"
}

func (this *UploadedSlicedFile) Put() ghost.Response {
	ctx := this.GetCtx()
	params := this.PutParams

	bs_media.NewUploader(ctx).UploadSlicedFile(&bs_media.SliceUploadParams{
		FileType:        db_file.FILE_TEXT2TYPE[params.FileType],
		GroupId:         params.GroupId,
		Filename:        params.Filename,
		FileHeader:      params.Slice,
		CompleteHash:    params.CompleteHash,
		SliceHash:       params.SliceHash,
		SliceIndex:      params.SliceIndex,
		TotalSliceCount: params.TotalSliceCount,
	})
	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&UploadedSlicedFile{})
}
