package file

import (
	"context"
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
)

type GroupUserFillService struct {
	ghost.DomainService
}

func (this *GroupUserFillService) FillUser(gUsers []*GroupUser) {
	userIds := make([]int, 0, len(gUsers))
	for _, gUser := range gUsers {
		userIds = append(userIds, gUser.UserId)
	}
	users := bm_account.NewUserRepository(this.GetCtx()).GetByIds(userIds)
	id2user := make(map[int]*bm_account.User)
	for _, user := range users {
		id2user[user.Id] = user
	}
	for _, gUser := range gUsers {
		gUser.User = id2user[gUser.UserId]
	}
}

func NewGroupUserFillService(ctx context.Context) *GroupUserFillService {
	inst := new(GroupUserFillService)
	inst.SetCtx(ctx)
	return inst
}
