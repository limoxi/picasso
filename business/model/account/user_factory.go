package account

import (
	"context"
	"github.com/limoxi/ghost"
	db_user "picasso/db/user"
)

type UserFactory struct {
	ghost.DomainService
}

func (this *UserFactory) Create(user *User) *User{
	dbModel := &db_user.User{
		Phone: user.Phone,
		Password: user.Password,
	}
	if err := ghost.GetDBFromCtx(this.GetCtx()).Create(&dbModel).Error; err != nil{
		panic(err)
	}
	ghost.Debug("[assert dbModel Id]", dbModel.Id)
	user.Id = dbModel.Id
	return user
}

func NewUserFactory(ctx context.Context) *UserFactory{
	inst := new(UserFactory)
	inst.SetCtx(ctx)
	return inst
}