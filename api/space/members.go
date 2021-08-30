package space

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_space "picasso/business/model/space"
)

type Members struct {
	ghost.ApiTemplate

	GetParams *struct{
		SpaceId int `form:"space_id"`
	}
}

func (this *Members) Resource() string{
	return "space.members"
}

func (this *Members) Get() ghost.Response{
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)
	space := m_space.NewSpaceRepository(ctx).GetById(this.GetParams.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	if !space.HasMember(user){
		panic(ghost.NewBusinessError("当前用户没有权限"))
	}

	members := space.GetMembers()
	return ghost.NewJsonResponse(ghost.Map{
		"members": m_space.NewSpaceEncodeService(ctx).EncodeManyMembers(members),
	})
}

func init(){
	ghost.RegisterApi(&Members{})
}
