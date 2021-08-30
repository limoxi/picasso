package media

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-cmd/cmd"
	"github.com/limoxi/ghost"
	db_media "picasso/db/media"
)

type MediaMetadataProcessor struct {
	ghost.DomainService
}

func (this *MediaMetadataProcessor) ProcessImage(dbModel *db_media.Media) *Metadata{
	result := <- cmd.NewCmd(
		"ExifTool",
		dbModel.StoragePath,
		"-c", "%.6f degrees",
		"-d", "%Y-%m-%d %H:%M:%S").Start()
	spew.Dump(result.Stdout)
	return nil
}

func (this *MediaMetadataProcessor) ProcessVideo() *Metadata{
	return nil
}

func NewMediaMetadataProcessor(ctx context.Context) *MediaMetadataProcessor{
	inst := new(MediaMetadataProcessor)
	inst.SetCtx(ctx)
	return inst
}