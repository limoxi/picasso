package file

import (
	"github.com/limoxi/ghost"
	bs_file "picasso/business/service/file"
)

type UploadedFileHashes struct {
	ghost.ApiTemplate

	PutParams *struct {
		Hashes []string `form:"hashes"`
	}
}

func (this *UploadedFileHashes) Resource() string {
	return "file.uploaded_file_hashes"
}

// Put 支持秒传，检查hashcode是否已存在
func (this *UploadedFileHashes) Put() ghost.Response {
	ctx := this.GetCtx()
	return ghost.NewJsonResponse(
		bs_file.NewFileService(ctx).CheckExistenceByHashes(this.PutParams.Hashes),
	)
}

func init() {
	ghost.RegisterApi(&UploadedFileHashes{})
}
