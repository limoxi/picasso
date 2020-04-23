package user

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	ds_account "picasso/domain/service/account"
)

type LoginedUser struct {
	ghost.ApiTemplate
}

func (this *LoginedUser) GetResource() string{
	return "user.logined_user"
}

func (this *LoginedUser) Put() ghost.Response{
	params := new(struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	})
	this.Bind(params)
	ctx := this.GetCtx()
	user := ds_account.NewLoginService(ctx).Login(params.Phone, params.Password)
	encodedUser := dm_account.NewUserEncodeService(ctx).Encode(user)
	encodedUser.Token = user.GetJWTToken()
	ctx.Header("Authorization", encodedUser.Token)
	return ghost.NewJsonResponse(encodedUser)
}

func init(){
	ghost.RegisterApi(&LoginedUser{})
}