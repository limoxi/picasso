package user

import (
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	app_account "picasso/business/app/account"
	m_account "picasso/business/model/account"
)

type LoginedUser struct {
	ghost.ApiTemplate

	PutParams *struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	}
}

func (this *LoginedUser) Resource() string{
	return "user.logined_user"
}

func (this *LoginedUser) Put() ghost.Response{
	ctx := this.GetCtx()
	user := app_account.NewLoginService(ctx).Login(this.PutParams.Phone, this.PutParams.Password)
	encodedUser := m_account.NewUserEncodeService(ctx).Encode(user)
	encodedUser.Token = user.GetJWTToken()
	ctx.(*gin.Context).Header("Authorization", encodedUser.Token)
	return ghost.NewJsonResponse(encodedUser)
}

func init(){
	ghost.RegisterApi(&LoginedUser{})
}