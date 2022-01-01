package file

import (
	"context"
	"github.com/limoxi/ghost"
	bm_account "picasso/business/model/account"
	db_file "picasso/db/file"
)

type GroupUser struct {
	ghost.DomainModel
	Id        int
	UserId    int
	NickName  string
	GroupId   int
	IsManager bool

	User *bm_account.User
}

func NewGroupUserFromDbModel(ctx context.Context, dbModel *db_file.GroupUser) *GroupUser {
	inst := new(GroupUser)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}
