package space

import (
	"context"
	"github.com/limoxi/ghost"
	db_space "picasso/db/space"
	dm_account "picasso/domain/model/account"
)

type SpaceMember struct {
	ghost.DomainModel
	Id int
	UserId int
	NickName string
	SpaceId int
	IsManager bool

	User *dm_account.User
}

func NewSpaceMemberFromDbModel(ctx context.Context, dbModel *db_space.SpaceMember) *SpaceMember{
	inst := new(SpaceMember)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}