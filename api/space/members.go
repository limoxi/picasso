package space

import (
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	bm_space "picasso/business/model/space"
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
	user := bm_account.GetUserFromCtx(ctx)
	space := bm_space.NewSpaceRepository(ctx).GetById(params.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	if !space.HasMember(user){
		panic(ghost.NewBusinessError("当前用户没有权限"))
	}

	members := space.GetMembers()
	return ghost.NewJsonResponse(ghost.Map{
		"members": bm_space.NewSpaceEncodeService(ctx).EncodeManyMembers(members),
	})
}

func init(){
	ghost.RegisterApi(&Members{})
}
