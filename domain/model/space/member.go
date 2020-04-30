package space

import (
	"context"
	"github.com/limoxi/ghost"
	m_space "picasso/common/db/space"
)

type Member struct {
	ghost.DomainModel
	Id int
	UserId int
	NickName string
	Code string
	SpaceId int
}


func NewSpaceMemberFromDbModel(ctx context.Context, dbModel *m_space.Space) *Member{
	inst := new(Member)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}