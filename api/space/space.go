package space

import (
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_space "picasso/business/model/space"
)

type Space struct {
	ghost.ApiTemplate

	PutParams *struct{
		Name string `json:"name"`
	}
}

func (this *Space) Resource() string{
	return "space.space"
}

func (this *Space) Put() ghost.Response{
	ctx := this.GetCtx()
	user := m_account.GetUserFromCtx(ctx)
	space := m_space.NewSpaceForUser(ctx, user, this.PutParams.Name)
	return ghost.NewJsonResponse(ghost.Map{
		"id": space.Id,
	})
}

func init(){
	ghost.RegisterApi(&Space{})
}
