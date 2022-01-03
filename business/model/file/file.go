package file

import (
	"context"
	"github.com/limoxi/ghost"
	"path"
	db_file "picasso/db/file"
	"time"
)

type File struct {
	ghost.DomainModel

	GroupId         int
	Type            int
	Hash            string
	Name            string
	StorageBasePath string
	StorageDirPath  string
	Status          int
	Metadata        string
	CreatedTime     time.Time
	ThumbnailPath   string
	Size            int64
}

func (this *File) GetFullPath() string {
	return path.Join(this.StorageBasePath, this.StorageDirPath, this.Name)
}

func NewFileFromDbModel(ctx context.Context, dbModel *db_file.File) *File {
	inst := new(File)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}
