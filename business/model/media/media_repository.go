package media

import (
	"context"
	"github.com/limoxi/ghost"
)

type MediaRepository struct {
	ghost.BasDomainRepository
}

func NewMediaRepository(ctx context.Context) *MediaRepository{
	inst := new(MediaRepository)
	inst.SetCtx(ctx)
	return inst
}