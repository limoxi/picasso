package user

import (
	"github.com/limoxi/ghost"
	ds_account "picasso/domain/service/account"
)

type RegisteredUser struct {
	ghost.ApiTemplate
}

func (this *RegisteredUser) GetResource() string{
	return "user.registered_user"
}

func (this *RegisteredUser) Put() ghost.Response{
	params := new(struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	})
	this.Bind(params)
	ctx := this.GetCtx()

	user := ds_account.NewLoginService(ctx).Register(ds_account.RegisterParams{
		Phone: params.Phone,
		RawPassword: params.Password,
	})
	return ghost.NewJsonResponse(ghost.Map{
		"id": user.Id,
		"token": user.GetJWTToken(),
	})
}

func init(){
	ghost.RegisterApi(&RegisteredUser{})
}