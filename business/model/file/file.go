package file

import (
	"context"
	"github.com/limoxi/ghost"
	db_file "picasso/db/file"
	"time"
)

type File struct {
	ghost.DomainModel

	GroupId       int
	Type          int
	SimpleHash    string
	FullHash      string
	ThumbnailPath string
	StoragePath   string
	Status        int
	Metadata      string
	ShootTime     time.Time
	ShootLocation string
	Size          int64
	Duration      int
}

func NewFileFromDbModel(ctx context.Context, dbModel *db_file.File) *File {
	inst := new(File)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}
