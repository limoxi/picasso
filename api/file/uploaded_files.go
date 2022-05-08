package file

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"mime/multipart"
	bm_account "picasso/business/model/account"
	bs_file "picasso/business/service/file"
)

type UploadedFiles struct {
	ghost.ApiTemplate

	PutParams *struct {
		Path   string                  `form:"path"`
		Files  []*multipart.FileHeader `form:"files"`
		Hashes string                  `form:"hashes"`
	}
}

// Resource 批量上传文件
func (this *UploadedFiles) Resource() string {
	return "file.uploaded_files"
}

func (this *UploadedFiles) Put() ghost.Response {
	ctx := this.GetCtx()
	params := this.PutParams

	var hashes []string
	if err := ghost_util.Decode(params.Hashes, &hashes); err != nil {
		panic(err)
	}

	bs_file.NewUploader(ctx).UploadFiles(&bs_file.UploadParams{
		User:        bm_account.GetUserFromCtx(ctx),
		Path:        params.Path,
		FileHeaders: params.Files,
		Hashes:      hashes,
	})
	return ghost.NewJsonResponse(ghost.Map{})
}

func init() {
	ghost.RegisterApi(&UploadedFiles{})
}
