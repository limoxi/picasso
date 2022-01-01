package file

import (
	"context"
	"github.com/limoxi/ghost"
	db_file "picasso/db/file"
)

type FileRepository struct {
	ghost.BaseDomainRepository
}

func (this *FileRepository) GetByFilters(filters ghost.Map) []*File {
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx).Model(&db_file.File{}).Where(filters).Order("-id")
	var dbModels []*db_file.File
	if this.Paginator != nil {
		this.Paginator.Paginate(db)
	}
	result := db.Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}

	medias := make([]*File, 0)
	for _, dbModel := range dbModels {
		medias = append(medias, NewFileFromDbModel(ctx, dbModel))
	}
	return medias
}

func (this *FileRepository) GetPagedFilesForUser(userId int, filters ghost.Map) []*File {
	if filters == nil {
		filters = ghost.Map{}
	}
	filters["user_id"] = userId
	return this.GetByFilters(filters)
}

func NewFileRepository(ctx context.Context) *FileRepository {
	inst := new(FileRepository)
	inst.SetCtx(ctx)
	return inst
}
