package media

import (
	"context"
	"github.com/limoxi/ghost"
	db_media "picasso/db/file"
)

type MediaRepository struct {
	ghost.BasDomainRepository
}

func (this *MediaRepository) GetByFilters(filters ghost.Map) []*Media {
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx).Model(&db_media.Media{}).Where(filters).Order("-id")
	var dbModels []*db_media.Media
	if this.Paginator != nil {
		this.Paginator.Paginate(db)
	}
	result := db.Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}

	medias := make([]*Media, 0)
	for _, dbModel := range dbModels {
		medias = append(medias, NewMediaFromDbModel(ctx, dbModel))
	}
	return medias
}

func (this *MediaRepository) GetPagedMediasForUser(userId int, filters ghost.Map) []*Media {
	if filters == nil {
		filters = ghost.Map{}
	}
	filters["user_id"] = userId
	return this.GetByFilters(filters)
}

func NewMediaRepository(ctx context.Context) *MediaRepository {
	inst := new(MediaRepository)
	inst.SetCtx(ctx)
	return inst
}
