package space

import (
	"context"
	"github.com/limoxi/ghost"
	m_space "picasso/common/db/space"
	dm_account "picasso/domain/model/account"
)

type SpaceRepository struct {
	ghost.DomainObject
}

func (this *SpaceRepository) GetByFilters(filters ghost.Map, paginator *ghost.Paginator) []*Space{
	db := ghost.GetDB().Model(&m_space.Space{}).Where(filters)
	var dbModels []*m_space.Space
	if paginator != nil{
		paginator.Paginate(db)
	}
	result := db.Order("id desc").Find(&dbModels)
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

func (this *SpaceRepository) GetSpacesForUser(user *dm_account.User, filters ghost.Map, paginator *ghost.Paginator) []*Space{
	filters["user_id"] = user.Id
	return this.GetByFilters(filters, paginator)
}

func (this *SpaceRepository) GetForUser(user *dm_account.User, spaceId int) *Space{
	spaces := this.GetByFilters(ghost.Map{
		"user_id": user.Id,
	}, nil)
	if len(spaces) > 0{
		return spaces[0]
	}
	return nil
}

func (this *SpaceRepository) GetById(spaceId int) *Space{
	spaces := this.GetByFilters(ghost.Map{
		"id": spaceId,
	}, nil)
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
