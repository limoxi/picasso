package space

import (
	"context"
	"github.com/limoxi/ghost"
	m_space "picasso/common/db/space"
	dm_account "picasso/domain/model/account"
)

type Space struct {
	ghost.DomainModel
	Id int
	Name string
}

func NewSpaceFromDbModel(ctx context.Context, dbModel *m_space.Space) *Space{
	inst := new(Space)
	inst.SetCtx(ctx)
	inst.Id = dbModel.Id
	inst.Name = dbModel.Name
	return inst
}

func NewSpaceForUser(ctx context.Context, user *dm_account.User, name string) *Space{
	dbModel := &m_space.Space{
		Name: name,
		UserId: user.Id,
	}
	result := ghost.GetDB().Create(dbModel)
	if err := result.Error; err != nil{
		panic(ghost.NewSystemError(err.Error(), "创建空间失败"))
	}
	return &Space{
		Id: dbModel.Id,
	}
}
