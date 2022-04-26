package upload

import (
	"github.com/limoxi/ghost"
	bs_file "picasso/business/service/file"
)

type UploadedFileSliceHashes struct {
	ghost.ApiTemplate

	PutParams *struct {
		CompleteHash string   `json:"complete_hash"`
		SliceHashes  []string `json:"slice_hashes"`
	}
}

func (this *UploadedFileSliceHashes) Resource() string {
	return "upload.uploaded_file_slice_hashes"
}

// Put 支持秒传
func (this *UploadedFileSliceHashes) Put() ghost.Response {
	ctx := this.GetCtx()
	params := this.PutParams
	return ghost.NewJsonResponse(
		bs_file.NewFileService(ctx).CheckSliceExistenceByHashes(
			params.CompleteHash, params.SliceHashes),
	)
}

func init() {
	ghost.RegisterApi(&UploadedFileSliceHashes{})
}
