package account

import (
	"context"
	"github.com/limoxi/ghost"
	db_user "picasso/db/user"
)

type UserFactory struct {
	ghost.DomainObject
}

// Create
func (*UserFactory) Create(user *User) *User{
	dbModel := &db_user.User{
		Phone: user.Phone,
		Password: user.Password,
	}
	if err := ghost.GetDB().Create(&dbModel).Error; err != nil{
		ghost.Panic(err)
	}
	user.Id = dbModel.Id
	return user
}

func NewUserFactory(ctx context.Context) *UserFactory{
	inst := new(UserFactory)
	inst.SetCtx(ctx)
	return inst
}