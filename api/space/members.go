package space

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type Members struct {
	ghost.ApiTemplate
}

type spaceMembersGetParams struct {
	SpaceId int `form:"space_id"`
}

func (this *Members) Resource() string{
	return "space.members"
}

func (this *Members) Get() ghost.Response{
	var params spaceMembersGetParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := dm_account.GetUserFromCtx(ctx)
	space := dm_space.NewSpaceRepository(ctx).GetById(params.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	if !space.HasMember(user){
		panic(ghost.NewBusinessError("当前用户没有权限"))
	}
	space.GetMembers()
	return ghost.NewJsonResponse(ghost.Map{})
}

func init(){
	ghost.RegisterApi(&Member{})
}
