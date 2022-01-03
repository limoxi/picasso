package file

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-cmd/cmd"
	"github.com/limoxi/ghost"
	bm_file "picasso/business/model/file"
	db_file "picasso/db/file"
)

type MediaMetadataProcessor struct {
	ghost.DomainService
}

func (this *MediaMetadataProcessor) getMediaType() int {
	return 0 // todo
}

func (this *MediaMetadataProcessor) Process(mediaFile *bm_file.File) *Metadata {
	mediaType := this.getMediaType()
	switch mediaType {
	case db_file.MEDIA_TYPE_IMAGE:
		return this.processImage(mediaFile)
	case db_file.MEDIA_TYPE_VIDEO:
		return this.processVideo()
	default:
		return nil
	}
}

func (this *MediaMetadataProcessor) processImage(mediaFile *bm_file.File) *Metadata {
	result := <-cmd.NewCmd(
		"ExifTool",
		mediaFile.GetFullPath(),
		"-c", "%.6f degrees",
		"-d", "%Y-%m-%d %H:%M:%S").Start()
	spew.Dump(result.Stdout)
	return nil
}

func (this *MediaMetadataProcessor) processVideo() *Metadata {
	return nil
}

func NewMediaMetadataProcessor(ctx context.Context) *MediaMetadataProcessor {
	inst := new(MediaMetadataProcessor)
	inst.SetCtx(ctx)
	return inst
}
