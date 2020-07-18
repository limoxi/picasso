package media

import (
	"context"
	"fmt"
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"io/ioutil"
	"os"
	"path"
	db_media "picasso/db/media"
	"strings"
)

type MediaService struct {
	ghost.DomainService
}

// CheckSliceExistenceByHashes
func (this *MediaService) CheckSliceExistenceByHashes(completeFilename, completeHash string, totalSliceCount int,
	sliceHashes []string, sliceHash2index map[string]int) map[string]bool{
	pureFileName := strings.Split(completeFilename, ".")[0]
	tmpDirPath := path.Join(VIDEO_STORAGE_PATH, fmt.Sprintf("tmp_%s_%s", pureFileName, completeHash))
	sliceHash2existed := make(map[string]bool)
	_, err := os.Stat(tmpDirPath)
	if err != nil{
		if os.IsNotExist(err){
			for _, sliceHash := range sliceHashes{
				sliceHash2existed[sliceHash] = false
			}
			return sliceHash2existed
		}else{
			ghost.Error(err)
			panic(err)
		}
	}
	fs, err := ioutil.ReadDir(tmpDirPath)
	if err != nil{
		ghost.Error(err)
		panic(err)
	}
	allFilenames := make([]string, 0, len(fs))
	for _, f := range fs{
		allFilenames = append(allFilenames, f.Name())
	}

	lister := ghost_util.NewListerFromStrings(allFilenames)
	for _, sliceHash := range sliceHashes{
		if lister.Has(fmt.Sprintf("%s_%s_%d_%d.slice",
			pureFileName, sliceHash, totalSliceCount, sliceHash2index[sliceHash])){
			sliceHash2existed[sliceHash] = true
		}else{
			sliceHash2existed[sliceHash] = false
		}
	}
	return sliceHash2existed
}

func (this *MediaService) CheckExistenceByHashes(hashes []string) map[string]bool{
	var dbModels []*db_media.Media
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_media.Media{}).Where(ghost.Map{
		"hash__in": hashes,
		"status__not": db_media.MEDIA_STATUS_SLICE_SAVED,
	}).Find(&dbModels)
	if err := result.Error; err != nil{
		ghost.Error(err)
		panic(err)
	}
	existedHashes := make([]string, 0, len(dbModels))
	for _, dbModel := range dbModels{
		existedHashes = append(existedHashes, dbModel.Hash)
	}
	lister := ghost_util.NewListerFromStrings(existedHashes)
	hash2existence := make(map[string]bool)
	for _, hash := range hashes{
		hash2existence[hash] = lister.Has(hash)
	}
	return hash2existence
}

func NewMediaService(ctx context.Context) *MediaService{
	inst := new(MediaService)
	inst.SetCtx(ctx)
	return inst
}
