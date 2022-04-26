package upload

import (
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"mime/multipart"
	bm_account "picasso/business/model/account"
	bs_file "picasso/business/service/file"
	db_file "picasso/db/file"
)

type UploadedFiles struct {
	ghost.ApiTemplate

	PutParams *struct {
		GroupId  int                     `form:"group_id"`
		FileType string                  `form:"file_type"`
		Files    []*multipart.FileHeader `form:"files"`
		Hashes   string                  `form:"hashes"`
	}
}

// Resource 批量上传文件
func (this *UploadedFiles) Resource() string {
	return "upload.uploaded_files"
}

func (this *UploadedFiles) Put() ghost.Response {
	ctx := this.GetCtx()
	params := this.PutParams
	fileType := db_file.FILE_TEXT2TYPE[params.FileType]
	if fileType == 0 {
		panic(ghost.NewBusinessError("不支持的文件类型"))
	}

	var hashes []string
	if err := ghost_util.Decode(params.Hashes, &hashes); err != nil {
		panic(err)
	}

	bs_file.NewUploader(ctx).UploadFiles(&bs_file.UploadParams{
		User:        bm_account.GetUserFromCtx(ctx),
		GroupId:     params.GroupId,
		FileType:    fileType,
		FileHeaders: params.Files,
		Hashes:      hashes,
	})
	return ghost.NewJsonResponse(ghost.Map{})
}

func init() {
	ghost.RegisterApi(&UploadedFiles{})
}
