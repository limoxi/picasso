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
		Id:               file.Id,
		Type:             file.Type,
		Hash:             file.Hash,
		Name:             file.Name,
		Path:             file.Path,
		Size:             file.Size,
		Status:           file.Status,
		Thumbnail:        file.Thumbnail,
		Metadata:         file.Metadata,
		CreatedTime:      file.CreatedTime.Format(common.DATETIME_LAYOUT),
		LastModifiedTime: file.LastModifiedTime.Format(common.DATETIME_LAYOUT),
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
