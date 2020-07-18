package space

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_space "picasso/business/model/space"
)

type Code struct {
	ghost.ApiTemplate
}

func (this *Code) Resource() string{
	return "space.code"
}

type spaceCodeParams struct {
	SpaceId int `json:"space_id"`
}

// Put 生成邀请码
func (this *Code) Put() ghost.Response{
	var params spaceCodeParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := bm_account.GetUserFromCtx(ctx)
	space := bm_space.NewSpaceRepository(ctx).GetForUser(user, params.SpaceId)
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
