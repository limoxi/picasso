package space

import (
	"context"
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
)

type SpaceMemberFillService struct {
	ghost.DomainService
}

func (this *SpaceMemberFillService) FillUser(members []*SpaceMember){
	userIds := make([]int, 0, len(members))
	for _, member := range members{
		userIds = append(userIds, member.UserId)
	}
	users := m_account.NewUserRepository(this.GetCtx()).GetByIds(userIds)
	id2user := make(map[int]*m_account.User)
	for _, user := range users{
		id2user[user.Id] = user
	}
	for _, member := range members{
		member.User = id2user[member.UserId]
	}
}

func NewSpaceMemberFillService(ctx context.Context) *SpaceMemberFillService{
	inst := new(SpaceMemberFillService)
	inst.SetCtx(ctx)
	return inst
}
