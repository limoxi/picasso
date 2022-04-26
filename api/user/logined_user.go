package user

import (
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	app_account "picasso/business/service/account"
)

type LoginedUser struct {
	ghost.ApiTemplate

	PutParams *struct {
		Phone    string `json:"phone" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
}

func (this *LoginedUser) Resource() string {
	return "user.logined_user"
}

func (this *LoginedUser) Put() ghost.Response {
	ctx := this.GetCtx()
	user := app_account.NewLoginService(ctx).Login(this.PutParams.Phone, this.PutParams.Password)
	encodedUser := m_account.NewUserEncodeService(ctx).Encode(user)
	encodedUser.Token = user.GetJWTToken()
	ctx.(*gin.Context).Header("Authorization", encodedUser.Token)
	return ghost.NewJsonResponse(encodedUser)
}

func init() {
	ghost.RegisterApi(&LoginedUser{})
}
