package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	bs_file "picasso/business/service/file"
)

type UploadedFileSliceHashes struct {
	ghost.ApiTemplate

	GetParams *struct {
		CompleteHash string `form:"complete_hash"`
		SLiceHashes  string `form:"slice_hashes"`
	}
}

func (this *UploadedFileSliceHashes) Resource() string {
	return "upload.uploaded_file_slice_hashes"
}

// Get 支持秒传
func (this *UploadedFileSliceHashes) Get() ghost.Response {
	ctx := this.GetCtx()
	params := this.GetParams
	var hashList []string
	if err := ghost_util.Decode(params.SLiceHashes, &hashList); err != nil {
		panic(err)
	}
	return ghost.NewJsonResponse(
		bs_file.NewFileService(ctx).CheckSliceExistenceByHashes(params.CompleteHash, hashList),
	)
}

func init() {
	ghost.RegisterApi(&UploadedFileSliceHashes{})
}
