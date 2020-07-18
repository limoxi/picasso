package user

import (
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bs_account "picasso/business/service/account"
)

type LoginedUser struct {
	ghost.ApiTemplate
}

func (this *LoginedUser) Resource() string{
	return "user.logined_user"
}

func (this *LoginedUser) Put() ghost.Response{
	params := new(struct{
		Phone string `json:"phone"`
		Password string `json:"password"`
	})
	this.Bind(params)
	ctx := this.GetCtx()
	user := bs_account.NewLoginService(ctx).Login(params.Phone, params.Password)
	encodedUser := bm_account.NewUserEncodeService(ctx).Encode(user)
	encodedUser.Token = user.GetJWTToken()
	ctx.(*gin.Context).Header("Authorization", encodedUser.Token)
	return ghost.NewJsonResponse(encodedUser)
}

func init(){
	ghost.RegisterApi(&LoginedUser{})
}