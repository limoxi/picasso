package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
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
}

func (this *Dir) Resource() string {
	return "file.dir"
}

func (this *Dir) Put() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.PutParams
	bm_file.NewDir(ctx, user, params.Path, params.Name)
	return ghost.NewJsonResponse(nil)
}

func (this *Dir) Post() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	params := this.PostParams

	file := bm_file.NewFileRepository(ctx).GetById(params.Id)
	if file == nil {
		panic(ghost.NewBusinessError("文件夹不存在"))
	}
	if file.UserId != user.Id {
		panic(ghost.NewBusinessError("当前用户无权限"))
	}
	file.UpdateName(params.Name)
	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&Dir{})
}
