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

func (this *Space) AddUser(user *dm_account.User){
	if this.HasUser(user){
		panic(ghost.NewBusinessError("已邀请用户或已存在"))
	}
	if err := ghost.GetDB().Create(&m_space.SpaceHasUser{
		SpaceId: this.Id,
		UserId: user.Id,
		IsManager: false,
		Status: m_space.SPACE_USER_STATUS_INVITED,
	}).Error; err != nil{
		ghost.Panic(err)
	}
}

func (this *Space) HasUser(user *dm_account.User) bool{
	filters := ghost.Map{
		"space_id": this.Id,
		"user_id": user.Id,
	}
	var count int
	result := ghost.GetDB().Model(&m_space.SpaceHasUser{}).Where(filters).Count(&count)
	if err := result.Error; err != nil{
		panic(err)
	}
	return count > 0
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
