package space

import (
	"context"
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
	db_space "picasso/db/space"
)

type SpaceMember struct {
	ghost.DomainModel
	Id int
	UserId int
	NickName string
	SpaceId int
	IsManager bool

	User *m_account.User
}

func NewSpaceMemberFromDbModel(ctx context.Context, dbModel *db_space.SpaceMember) *SpaceMember{
	inst := new(SpaceMember)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}