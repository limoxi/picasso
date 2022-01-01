package file

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
)

type Group struct {
	ghost.ApiTemplate

	PutParams *struct {
		Name string `json:"name"`
	}
}

func (this *Group) Resource() string {
	return "user.group"
}

func (this *Group) Put() ghost.Response {
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	group := bm_file.NewGroupForUser(ctx, user, this.PutParams.Name)
	return ghost.NewJsonResponse(ghost.Map{
		"id": group.Id,
	})
}

func init() {
	ghost.RegisterApi(&Group{})
}
