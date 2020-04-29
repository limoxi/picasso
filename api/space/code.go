package space

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type Code struct {
	ghost.ApiTemplate
}

func (this *Code) Resource() string{
	return "space.code"
}

type spaceCodeParams struct {
	SpaceId int `form:"space_id"`
}

// Put 生成邀请码
func (this *Code) Put() ghost.Response{
	var params spaceCodeParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := dm_account.GetUserFromCtx(ctx)
	space := dm_space.NewSpaceRepository(ctx).GetForUser(user, params.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	return ghost.NewJsonResponse(ghost.Map{
		"code": space.GenCode(),
	})
}

func init(){
	ghost.RegisterApi(&Code{})
}
