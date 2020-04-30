package space

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type Space struct {
	ghost.ApiTemplate
}

type spacePutParams struct {
	Name string `json:"name"`
}

func (this *Space) Resource() string{
	return "space.space"
}

func (this *Space) Put() ghost.Response{
	var params spacePutParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := dm_account.GetUserFromCtx(ctx)
	space := dm_space.NewSpaceForUser(ctx, user, params.Name)
	return ghost.NewJsonResponse(ghost.Map{
		"id": space.Id,
	})
}

func init(){
	ghost.RegisterApi(&Space{})
}
