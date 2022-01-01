package file

import (
	"context"
	"github.com/limoxi/ghost"
	"picasso/common"
)

type FileEncodeService struct {
	ghost.DomainService
}

func (this *FileEncodeService) Encode(file *File) *EncodedFile {
	return &EncodedFile{
		GroupId:       file.GroupId,
		Type:          file.Type,
		SimpleHash:    file.SimpleHash,
		FullHash:      file.FullHash,
		ThumbnailPath: file.ThumbnailPath,
		StoragePath:   file.StoragePath,
		Status:        file.Status,
		Metadata:      file.Metadata,
		ShootTime:     file.ShootTime.Format(common.DATETIME_LAYOUT),
		ShootLocation: file.ShootLocation,
		Size:          file.Size,
		Duration:      file.Duration,
	}
}

func (this *FileEncodeService) EncodeMany(files []*File) []*EncodedFile {
	encodedRecords := make([]*EncodedFile, 0, len(files))
	for _, file := range files {
		encodedRecords = append(encodedRecords, this.Encode(file))
	}
	return encodedRecords
}

func NewFileEncodeService(ctx context.Context) *FileEncodeService {
	inst := new(FileEncodeService)
	inst.SetCtx(ctx)
	return inst
}
