package file

import (
	"context"
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"os"
	"path"
	bm_account "picasso/business/model/account"
	bm_file "picasso/business/model/file"
	db_file "picasso/db/file"
	"strings"
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

func (this *FileService) AddDir(user *bm_account.User, dirPath, dirName string) {
	db := ghost.GetDBFromCtx(this.GetCtx())
	if db.Model(&db_file.File{}).Where(ghost.Map{
		"user_id": user.Id,
		"type":    db_file.FILE_TYPE_DIR,
		"path":    dirPath,
		"name":    dirName,
	}).Exist() {
		panic(ghost.NewBusinessError("目录已存在"))
	}
	result := db.Create(&db_file.File{
		UserId: user.Id,
		Type:   db_file.FILE_TYPE_DIR,
		Path:   dirPath,
		Status: db_file.FILE_STATUS_COMPLETE,
		Name:   dirName,
	})
	if result.Error != nil {
		ghost.Error(result.Error)
		panic(ghost.NewSystemError("添加目录失败"))
	}

	realPath := strings.ReplaceAll(path.Join(dirPath, dirName), ".", string(os.PathSeparator))
	ghost.Info(realPath, "===========")
	targetPath := path.Join(STORAGE_ROOT_PATH, realPath)
	err := os.Mkdir(targetPath, os.ModeDir)
	if err != nil {
		ghost.Error(err)
		panic(err)
	}
}

func (this *FileService) ChangeDirName(user *bm_account.User, dirId int, dirName string) {
	file := bm_file.NewFileRepository(this.GetCtx()).GetById(dirId)
	if file == nil {
		panic(ghost.NewBusinessError("文件夹不存在"))
	}
	if file.UserId != user.Id {
		panic(ghost.NewBusinessError("当前用户无权限"))
	}
	file.UpdateName(dirName)
}

func (this *FileService) DeleteDir(user *bm_account.User, dirId int, dirPath string) {

}

// MakeThumbnail 生成文件缩略图
func (this *FileService) MakeThumbnail(f *bm_file.File) {

}

func (this *FileService) EncryptDir(user *bm_account.User, dirId int) {

}

func (this *FileService) DecryptDir(user *bm_account.User, dirId int) {

}

func NewFileService(ctx context.Context) *FileService {
	inst := new(FileService)
	inst.SetCtx(ctx)
	return inst
}
