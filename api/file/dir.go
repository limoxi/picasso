package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bs_file "picasso/business/service/file"
)

type Dir struct {
	ghost.ApiTemplate

	PutParams *struct {
		Path string `json:"path"`
		Name string `json:"name"`
	}

	PostParams *struct {
		Id   int    `json:"int"`
		Name string `json:"name"`
	}

	DeleteParams *struct {
		Id   int    `json:"int"`
		Path string `json:"path"`
	}
}

func (this *Dir) Resource() string {
	return "file.dir"
}

func (this *Dir) Put() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.PutParams
	bs_file.NewFileService(ctx).AddDir(user, params.Path, params.Name)
	return ghost.NewJsonResponse(nil)
}

func (this *Dir) Post() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.PostParams

	bs_file.NewFileService(ctx).ChangeDirName(user, params.Id, params.Name)

	return ghost.NewJsonResponse(nil)
}

func (this *Dir) Delete() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.DeleteParams

	bs_file.NewFileService(ctx).DeleteDir(user, params.Id, params.Path)

	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&Dir{})
}
