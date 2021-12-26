package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	bs_media "picasso/business/app/file"
)

type UploadedFileHashes struct {
	ghost.ApiTemplate

	GetParams *struct {
		Hashes string `form:"hashes"`
	}
}

func (this *UploadedFileHashes) Resource() string {
	return "upload.uploaded_file_hashes"
}

// Get 支持秒传，检查hashcode是否已存在
func (this *UploadedFileHashes) Get() ghost.Response {
	ctx := this.GetCtx()
	var hashList []string
	if err := ghost_util.Decode(this.GetParams.Hashes, &hashList); err != nil {
		panic(err)
	}
	return ghost.NewJsonResponse(
		bs_media.NewFileService(ctx).CheckExistenceByHashes(hashList),
	)
}

func init() {
	ghost.RegisterApi(&UploadedFileHashes{})
}
