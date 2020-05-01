package space

import (
	"context"
	"github.com/limoxi/ghost"
	db_space "picasso/db/space"
	dm_account "picasso/domain/model/account"
)

type SpaceRepository struct {
	ghost.BasDomainRepository
}

func (this *SpaceRepository) GetByFilters(filters ghost.Map) []*Space{
	db := ghost.GetDB().Model(&db_space.Space{}).Where(filters)
	var dbModels []*db_space.Space
	if this.Paginator != nil{
		this.Paginator.Paginate(db)
	}
	result := db.Order("-id").Find(&dbModels)
	if err := result.Error; err != nil{
		panic(err)
	}
	ctx := this.GetCtx()
	spaces := make([]*Space, 0)
	for _, dbModel := range dbModels{
		spaces = append(spaces, NewSpaceFromDbModel(ctx, dbModel))
	}
	return spaces
}

func (this *SpaceRepository) GetSpacesForUser(user *dm_account.User, filters ghost.Map) []*Space{
	var dbModels []*db_space.SpaceMember
	result := ghost.GetDB().Model(&db_space.SpaceMember{}).Where(ghost.Map{
		"user_id": user.Id,
	}).Find(&dbModels)
	if err := result.Error; err != nil{
		panic(err)
	}
	spaceIds := make([]int, 0, len(dbModels))
	for _, dbModel := range dbModels{
		spaceIds = append(spaceIds, dbModel.SpaceId)
	}
	return this.GetByFilters(ghost.Map{
		"id__in": spaceIds,
	})
}

func (this *SpaceRepository) GetForUser(user *dm_account.User, spaceId int) *Space{
	spaces := this.GetByFilters(ghost.Map{
		"id": spaceId,
		"user_id": user.Id,
	})
	if len(spaces) > 0{
		return spaces[0]
	}
	return nil
}

func (this *SpaceRepository) GetById(spaceId int) *Space{
	spaces := this.GetByFilters(ghost.Map{
		"id": spaceId,
	})
	if len(spaces) > 0{
		return spaces[0]
	}
	return nil
}

func NewSpaceRepository(ctx context.Context) *SpaceRepository{
	inst := new(SpaceRepository)
	inst.SetCtx(ctx)
	return inst
}
