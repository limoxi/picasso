package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	ghost_utils "github.com/limoxi/ghost/utils"
	"picasso/common/util"
	dm_account "picasso/domain/model/account"
)

type AccountMiddleware struct{

}

func (this *AccountMiddleware) Init(){
	ghost.Info("AccountMiddleware loaded")
}

func (this *AccountMiddleware)  (ctx *gin.Context){
	if ghost_utils.NewListerFromStrings([]string{"/user/logined_user/"}).Has(ctx.FullPath()){
		return
	}
	token := ctx.GetHeader("Authorization")
	ghost.Debug(token, "+==========")
	if token != ""{
		data := util.DecodeJwtToken(token)
		userId := 0
		switch data["user_id"].(type) {
		case float64:
			userId = int(data["user_id"].(float64))
		case int:
			userId = data["user_id"].(int)
		}
		ctx.Set("user", dm_account.NewUserFromId(ctx, userId))
	}
}

func (this *AccountMiddleware) AfterResponse(ctx *gin.Context){

}

func init(){
	ghost.RegisterMiddleware(&AccountMiddleware{})
}