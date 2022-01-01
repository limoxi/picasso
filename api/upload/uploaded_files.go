package upload

import (
	"github.com/limoxi/ghost"
	"mime/multipart"
	bs_media "picasso/business/app/file"
	m_account "picasso/business/model/account"
	db_file "picasso/db/file"
)

type UploadedFiles struct {
	ghost.ApiTemplate

	// filename格式必须为：fileType.hash.originFileName
	PutParams *struct {
		GroupId  int                     `form:"group_id"`
		FileType string                  `form:"file_type"`
		Files    []*multipart.FileHeader `form:"files"`
	}
}

// Resource 批量上传文件
func (this *UploadedFiles) Resource() string {
	return "upload.uploaded_files"
}

func (this *UploadedFiles) Put() ghost.Response {
	ctx := this.GetCtx()
	params := this.PutParams

	bs_media.NewUploader(ctx).UploadFiles(&bs_media.UploadParams{
		User:        m_account.GetUserFromCtx(ctx),
		GroupId:     params.GroupId,
		FileType:    db_file.FILE_TEXT2TYPE[params.FileType],
		FileHeaders: params.Files,
	})
	return ghost.NewJsonResponse(ghost.Map{})
}

func init() {
	ghost.RegisterApi(&UploadedFiles{})
}
