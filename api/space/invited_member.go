package space

import (
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type InvitedMember struct {
	ghost.ApiTemplate
}

type invitedMemberPutParams struct {
	UserId int `form:"user_id"`
	SpaceId int `form:"space_id"`
}

func (this *InvitedMember) GetResource() string{
	return "space.invited_member"
}

func (this *InvitedMember) Put() ghost.Response{
	var params invitedMemberPutParams
	this.Bind(&params)
	ctx := this.GetCtx()
	user := dm_account.GetUserFromCtx(ctx)
	space := dm_space.NewSpaceRepository(ctx).GetForUser(user, params.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	space.AddUser(dm_account.NewUserFromId(ctx, params.SpaceId))
	return ghost.NewJsonResponse(ghost.Map{})
}

func init(){
	ghost.RegisterApi(&InvitedMember{})
}
