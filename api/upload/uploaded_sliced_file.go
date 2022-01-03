package upload

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	bs_file "picasso/business/service/file"
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
		TotalSliceCount int                   `form:"total_slice_count"`
		Slice           *multipart.FileHeader `form:"slice"`
	}
}

func (this *UploadedSlicedFile) Resource() string {
	return "upload.uploaded_sliced_file"
}

func (this *UploadedSlicedFile) Put() ghost.Response {
	ctx := this.GetCtx()
	params := this.PutParams

	bs_file.NewUploader(ctx).UploadSlicedFile(&bs_file.SliceUploadParams{
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
