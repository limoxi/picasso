package space

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_space "picasso/business/model/space"
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
	user := bm_account.GetUserFromCtx(ctx)
	space := bm_space.NewSpaceForUser(ctx, user, params.Name)
	return ghost.NewJsonResponse(ghost.Map{
		"id": space.Id,
	})
}

func init(){
	ghost.RegisterApi(&Space{})
}
