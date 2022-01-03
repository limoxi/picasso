package file

import (
	"context"
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	db_file "picasso/db/file"
)

type FileService struct {
	ghost.DomainService
}

func (this *FileService) CheckSliceExistenceByHashes(completeHash string, sliceHashes []string) map[string]bool {
	var dbModels []*db_file.SlicedFile
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_file.SlicedFile{}).Where(ghost.Map{
		"file_hash":      completeHash,
		"slice_hash__in": sliceHashes,
	}).Find(&dbModels)
	if err := result.Error; err != nil {
		ghost.Error(err)
		panic(err)
	}
	hash2existed := make(map[string]bool)
	for _, dbModel := range dbModels {
		hash2existed[dbModel.SliceHash] = true
	}
	hash2existence := make(map[string]bool)
	for _, sh := range sliceHashes {
		hash2existence[sh] = false
		if _, ok := hash2existed[sh]; ok {
			hash2existence[sh] = true
		}
	}
	return hash2existence
}

func (this *FileService) CheckExistenceByHashes(hashes []string) map[string]bool {
	var dbModels []*db_file.File
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_file.File{}).Where(ghost.Map{
		"hash__in":    hashes,
		"status__not": db_file.FILE_STATUS_SLICE_SAVED,
	}).Find(&dbModels)
	if err := result.Error; err != nil {
		ghost.Error(err)
		panic(err)
	}
	existedHashes := make([]string, 0, len(dbModels))
	for _, dbModel := range dbModels {
		existedHashes = append(existedHashes, dbModel.Hash)
	}
	lister := ghost_util.NewListerFromStrings(existedHashes)
	hash2existence := make(map[string]bool)
	for _, hash := range hashes {
		hash2existence[hash] = lister.Has(hash)
	}
	return hash2existence
}

func NewFileService(ctx context.Context) *FileService {
	inst := new(FileService)
	inst.SetCtx(ctx)
	return inst
}
