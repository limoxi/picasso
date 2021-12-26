package media

import (
	"context"
	"github.com/limoxi/ghost"
	db_media "picasso/db/file"
	"time"
)

type Media struct {
	ghost.DomainModel

	SpaceId       int
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

func NewMediaFromDbModel(ctx context.Context, dbModel *db_media.Media) *Media {
	inst := new(Media)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}
