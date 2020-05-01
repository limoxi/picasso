package space

import (
	"context"
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
)

type SpaceMemberFillService struct {
	ghost.DomainObject
}

func (this *SpaceMemberFillService) FillUser(members []*SpaceMember){
	userIds := make([]int, 0, len(members))
	for _, member := range members{
		userIds = append(userIds, member.UserId)
	}
	users := dm_account.NewUserRepository(this.GetCtx()).GetByIds(userIds)
	id2user := make(map[int]*dm_account.User)
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
