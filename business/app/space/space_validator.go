package space

import (
	"context"
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	m_space "picasso/business/model/space"
)

type SpaceValidator struct {
	ghost.DomainService
}

func (this *SpaceValidator) Check(user *m_account.User, spaceId int) (*m_space.Space, error){
	space := m_space.NewSpaceRepository(this.GetCtx()).GetById(spaceId)
	if space != nil{
		return nil, ghost.NewBusinessError("空间不存在")
	}
	if !space.HasMember(user){
		return nil, ghost.NewBusinessError("该用户无权限")
	}
	return space, nil
}

func NewSpaceValidator(ctx context.Context) *SpaceValidator{
	inst := new(SpaceValidator)
	inst.SetCtx(ctx)
	return inst
}