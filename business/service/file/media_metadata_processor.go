package file

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-cmd/cmd"
	"github.com/limoxi/ghost"
	bm_file "picasso/business/model/file"
)

type MediaMetadataProcessor struct {
	ghost.DomainService
}

func (this *MediaMetadataProcessor) getMediaType() string {
	return "" // todo
}

func (this *MediaMetadataProcessor) Process(mediaFile *bm_file.File) *Metadata {
	switch this.getMediaType() {
	case "image":
		return this.processImage(mediaFile)
	case "video":
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
