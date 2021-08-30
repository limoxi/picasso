package space

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_space "picasso/business/model/space"
)

type Code struct {
	ghost.ApiTemplate

	PutParams *struct{
		SpaceId int `json:"space_id"`
	}
}

// Resource 邀请码
func (this *Code) Resource() string{
	return "space.code"
}

func (this *Code) Put() ghost.Response{
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)
	space := m_space.NewSpaceRepository(ctx).GetForUser(user, this.PutParams.SpaceId)
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
