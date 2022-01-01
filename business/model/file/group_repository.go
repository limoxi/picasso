package file

import (
	"context"
	"github.com/limoxi/ghost"
	"picasso/business/model/account"
	db_file "picasso/db/file"
)

type GroupRepository struct {
	ghost.BaseDomainRepository
}

func (this *GroupRepository) GetByFilters(filters ghost.Map) []*Group {
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx).Model(&db_file.Group{}).Where(filters)
	var dbModels []*db_file.Group
	if this.Paginator != nil {
		this.Paginator.Paginate(db)
	}
	result := db.Order("-id").Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}

	groups := make([]*Group, 0)
	for _, dbModel := range dbModels {
		groups = append(groups, NewGroupFromDbModel(ctx, dbModel))
	}
	return groups
}

func (this *GroupRepository) GetGroupsForUser(user *account.User, filters ghost.Map) []*Group {
	var dbModels []*db_file.GroupUser
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_file.GroupUser{}).Where(ghost.Map{
		"user_id": user.Id,
	}).Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}
	groupIds := make([]int, 0, len(dbModels))
	for _, dbModel := range dbModels {
		groupIds = append(groupIds, dbModel.GroupId)
	}
	return this.GetByFilters(ghost.Map{
		"id__in": groupIds,
	})
}

func (this *GroupRepository) GetForUser(user *account.User, spaceId int) *Group {
	groups := this.GetByFilters(ghost.Map{
		"id":      spaceId,
		"user_id": user.Id,
	})
	if len(groups) > 0 {
		return groups[0]
	}
	return nil
}

func (this *GroupRepository) GetDefaultForUser(user *account.User) *Group {
	groups := this.GetGroupsForUser(user, ghost.Map{})
	if len(groups) > 0 {
		return groups[0]
	}
	dbModel := db_file.Group{
		Name:   "default",
		UserId: user.Id,
	}
	if err := ghost.GetDBFromCtx(this.GetCtx()).Create(&dbModel).Error; err != nil {
		panic(err)
	}
	return NewGroupFromDbModel(this.GetCtx(), &dbModel)
}

func (this *GroupRepository) GetById(spaceId int) *Group {
	groups := this.GetByFilters(ghost.Map{
		"id": spaceId,
	})
	if len(groups) > 0 {
		return groups[0]
	}
	return nil
}

func NewGroupRepository(ctx context.Context) *GroupRepository {
	inst := new(GroupRepository)
	inst.SetCtx(ctx)
	return inst
}
