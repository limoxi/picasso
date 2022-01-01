package account

import (
	"context"
	"github.com/limoxi/ghost"
	db_user "picasso/db/user"
)

type UserRepository struct {
	ghost.DomainService
}

// UserExisted 用户是否存在
func (this *UserRepository) UserExisted(phone string) bool {
	var count int64
	filters := ghost.Map{
		"phone": phone,
	}
	ghost.GetDBFromCtx(this.GetCtx()).Model(&db_user.User{}).Where(filters).Count(&count)
	return count > 0
}

func (this *UserRepository) GetByFilters(filters ghost.Map) []*User {
	var dbModels []*db_user.User
	users := make([]*User, 0)
	ctx := this.GetCtx()
	result := ghost.GetDBFromCtx(ctx).Where(filters).Find(&dbModels)
	if err := result.Error; err != nil {
		ghost.Error(err.Error())
		return users
	}
	for _, dbModel := range dbModels {
		users = append(users, NewUserFromDbModel(ctx, dbModel))
	}
	return users
}

func (this *UserRepository) GetByPhone(phone string) *User {
	filters := ghost.Map{
		"phone": phone,
	}
	users := this.GetByFilters(filters)
	if len(users) > 0 {
		return users[0]
	}
	return nil
}

func (this *UserRepository) GetById(uid int) *User {
	filters := ghost.Map{
		"id": uid,
	}
	users := this.GetByFilters(filters)
	if len(users) > 0 {
		return users[0]
	}
	return nil
}

func (this *UserRepository) GetByIds(userIds []int) []*User {
	filters := ghost.Map{
		"id__in": userIds,
	}
	return this.GetByFilters(filters)
}

func NewUserRepository(ctx context.Context) *UserRepository {
	inst := new(UserRepository)
	inst.SetCtx(ctx)
	return inst
}
