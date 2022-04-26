package user

import (
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	app_account "picasso/business/service/account"
)

type RegisteredUser struct {
	ghost.ApiTemplate

	PutParams *struct {
		Avatar   string `json:"avatar"`
		Nickname string `json:"nickname"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
}

func (this *RegisteredUser) Resource() string {
	return "user.registered_user"
}

func (this *RegisteredUser) Put() ghost.Response {
	ctx := this.GetCtx()
	user := app_account.NewLoginService(ctx).Register(app_account.RegisterParams{
		Avatar:      this.PutParams.Avatar,
		Nickname:    this.PutParams.Nickname,
		Phone:       this.PutParams.Phone,
		RawPassword: this.PutParams.Password,
	})
	encodedUser := m_account.NewUserEncodeService(ctx).Encode(user)
	encodedUser.Token = user.GetJWTToken()
	ctx.(*gin.Context).Header("Authorization", encodedUser.Token)
	return ghost.NewJsonResponse(encodedUser)
}

func init() {
	ghost.RegisterApi(&RegisteredUser{})
}
