package file

import (
	"context"
	"github.com/limoxi/ghost"
	db_file "picasso/db/file"
)

type GroupUserRepository struct {
	ghost.BaseDomainRepository
}

func (this *GroupUserRepository) GetByFilters(filters ghost.Map) []*GroupUser {
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx).Model(&db_file.GroupUser{}).Where(filters)
	var dbModels []*db_file.GroupUser
	if this.Paginator != nil {
		this.Paginator.Paginate(db)
	}
	result := db.Order("id").Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}

	gUsers := make([]*GroupUser, 0)
	for _, dbModel := range dbModels {
		gUsers = append(gUsers, NewGroupUserFromDbModel(ctx, dbModel))
	}
	return gUsers
}

func (this *GroupUserRepository) GetById(groupId, userId int) *GroupUser {
	gUsers := this.GetByFilters(ghost.Map{
		"group_id": groupId,
		"user_id":  userId,
	})
	if len(gUsers) > 0 {
		return gUsers[0]
	}
	return nil
}

func NewGroupUserRepository(ctx context.Context) *GroupUserRepository {
	inst := new(GroupUserRepository)
	inst.SetCtx(ctx)
	return inst
}
