package user

import (
	"github.com/limoxi/ghost"
	app_account "picasso/business/app/account"
)

type RegisteredUser struct {
	ghost.ApiTemplate

	PutParams *struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	}
}

func (this *RegisteredUser) Resource() string{
	return "user.registered_user"
}

func (this *RegisteredUser) Put() ghost.Response{
	ctx := this.GetCtx()
	user := app_account.NewLoginService(ctx).Register(app_account.RegisterParams{
		Phone: this.PutParams.Phone,
		RawPassword: this.PutParams.Password,
	})
	return ghost.NewJsonResponse(ghost.Map{
		"id": user.Id,
		"token": user.GetJWTToken(),
	})
}

func init(){
	ghost.RegisterApi(&RegisteredUser{})
}