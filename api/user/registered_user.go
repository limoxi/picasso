package user

import (
	"github.com/limoxi/ghost"
	bs_account "picasso/business/service/account"
)

type RegisteredUser struct {
	ghost.ApiTemplate
}

func (this *RegisteredUser) Resource() string{
	return "user.registered_user"
}

func (this *RegisteredUser) Put() ghost.Response{
	params := new(struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	})
	this.Bind(params)
	ctx := this.GetCtx()

	user := bs_account.NewLoginService(ctx).Register(bs_account.RegisterParams{
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