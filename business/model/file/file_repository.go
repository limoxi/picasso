package file

import (
	"context"
	"github.com/limoxi/ghost"
	db_file "picasso/db/file"
)

type FileRepository struct {
	ghost.BaseDomainRepository
}

func (this *FileRepository) GetByFilters(filters ghost.Map, orderAttrs ...string) []*File {
	ctx := this.GetCtx()
	db := ghost.GetDBFromCtx(ctx).Model(&db_file.File{}).Where(filters)
	if len(orderAttrs) == 0 {
		orderAttrs = append(orderAttrs, "-id")
	}
	db = db.Order(orderAttrs)
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

func (this *FileRepository) GetOrderedFilesForUser(userId int, filters ghost.Map, orderAttrs []string) []*File {
	filters["user_id"] = userId
	filters["type"] = db_file.FILE_TYPE_DIR
	dirs := this.GetByFilters(filters, orderAttrs...)
	filters["type"] = db_file.FILE_TYPE_FILE
	files := this.GetByFilters(filters, orderAttrs...)
	dirs = append(dirs, files...)
	return dirs
}

func (this *FileRepository) GetById(id int) *File {
	files := this.GetByFilters(ghost.Map{
		"id": id,
	})
	if len(files) > 0 {
		return files[0]
	}
	return nil
}

func NewFileRepository(ctx context.Context) *FileRepository {
	inst := new(FileRepository)
	inst.SetCtx(ctx)
	return inst
}
