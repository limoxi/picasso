package account

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/limoxi/ghost"
	"picasso/common/util"
	db_user "picasso/db/user"
	"time"
)

type User struct{
	ghost.DomainModel

	Id int
	Phone string
	Password string
	CreatedAt time.Time
}

func (this *User) Update(nickname, avatar string){
	result := ghost.GetDB().Model(&db_user.User{}).Where("id=?", this.Id).Update(ghost.Map{
		"nickname": nickname,
		"avatar": avatar,
	})
	if err := result.Error; err != nil{
		panic(ghost.NewSystemError(err.Error(), "更新用户信息失败"))
	}
}

func (this *User) UpdatePassword(oldPwd, newPwd string) {

}

func NewUserFromDbModel(ctx context.Context, dbModel *db_user.User) *User{
	inst := &User{}
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}

func NewUserFromId(ctx context.Context, id int) *User{
	inst := &User{}
	inst.Id = id
	inst.SetCtx(ctx)
	return inst
}

func GetUserFromCtx(ctx context.Context) *User{
	iuser, ok := ctx.(*gin.Context).Get("user")
	if ok{
		return iuser.(*User)
	}
	return nil
}

// 获取jwt_token
func (this *User) GetJWTToken() string{
	return util.EncodeJwtToken(ghost.Map{
		"user_id": this.Id,
		"phone": this.Phone,
	}, time.Hour * 24)
}