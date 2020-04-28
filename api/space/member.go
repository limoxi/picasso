package space

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type Member struct {
	ghost.ApiTemplate
}

type spaceMemberPutParams struct {
	SpaceId int `form:"space_id"`
	Code string `form:"code"`
}

func (this *Member) GetResource() string{
	return "space.member"
}

func (this *Member) Put() ghost.Response{
	var params spaceMemberPutParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := dm_account.GetUserFromCtx(ctx)
	space := dm_space.NewSpaceRepository(ctx).GetById(params.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	space.AddUser(user, params.Code)
	return ghost.NewJsonResponse(ghost.Map{})
}

func init(){
	ghost.RegisterApi(&Member{})
}
