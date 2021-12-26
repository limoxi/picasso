package media

import (
	"context"
	"github.com/limoxi/ghost"
	"picasso/common"
)

type MediaEncodeService struct {
	ghost.DomainService
}

func (this *MediaEncodeService) Encode(media *Media) *EncodedMedia{
	return &EncodedMedia{
		SpaceId:       media.SpaceId,
		Type:          media.Type,
		SimpleHash:    media.SimpleHash,
		FullHash:      media.FullHash,
		ThumbnailPath: media.ThumbnailPath,
		StoragePath:   media.StoragePath,
		Status:        media.Status,
		Metadata:      media.Metadata,
		ShootTime:     media.ShootTime.Format(common.DATETIME_LAYOUT),
		ShootLocation: media.ShootLocation,
		Size:          media.Size,
		Duration:      media.Duration,
	}
}

func (this *MediaEncodeService) EncodeMany(medias []*Media) []*EncodedMedia{
	encodedRecords := make([]*EncodedMedia, 0, len(medias))
	for _, media := range medias{
		encodedRecords = append(encodedRecords, this.Encode(media))
	}
	return encodedRecords
}

func NewMediaEncodeService(ctx context.Context) *MediaEncodeService{
	inst := new(MediaEncodeService)
	inst.SetCtx(ctx)
	return inst
}