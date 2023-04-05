package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bs_file "picasso/business/service/file"
)

type EncryptedDir struct {
	ghost.ApiTemplate

	PutParams *struct {
		Id int `json:"int"`
	}

	DeleteParams *struct {
		Id int `json:"int"`
	}
}

// Resource 加密的目录
func (this *EncryptedDir) Resource() string {
	return "file.encrypted_dir"
}

func (this *EncryptedDir) Put() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.PutParams
	bs_file.NewFileService(ctx).EncryptDir(user, params.Id)
	return ghost.NewJsonResponse(nil)
}

func (this *EncryptedDir) Delete() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.DeleteParams

	bs_file.NewFileService(ctx).DecryptDir(user, params.Id)

	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&EncryptedDir{})
}
