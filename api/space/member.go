package space

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_space "picasso/business/model/space"
)

type Member struct {
	ghost.ApiTemplate

	PutParams *struct{
		SpaceId int `json:"space_id"`
		Code string `json:"code"`
	}
}

func (this *Member) Resource() string{
	return "space.member"
}

func (this *Member) Put() ghost.Response{
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)
	space := m_space.NewSpaceRepository(ctx).GetById(this.PutParams.SpaceId)
	if space == nil{
		panic(ghost.NewBusinessError("空间不存在"))
	}
	space.AddMember(user, this.PutParams.Code)
	return ghost.NewJsonResponse(ghost.Map{})
}

func init(){
	ghost.RegisterApi(&Member{})
}
