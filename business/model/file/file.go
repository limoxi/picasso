package file

import (
	"context"
	"github.com/limoxi/ghost"
	"path"
	bm_account "picasso/business/model/account"
	db_file "picasso/db/file"
	"time"
)

type File struct {
	ghost.DomainModel

	Id     int
	UserId int
	Type   int
	Hash   string
	Path   string
	Status int

	Name             string
	Size             int64
	Metadata         string
	Thumbnail        string
	LastModifiedTime time.Time
	CreatedTime      time.Time
}

func (this *File) GetFullPath() string {
	return path.Join(this.Path, this.Name)
}

func (this *File) UpdateName(name string) {
	if name == this.Name {
		return
	}
	db := ghost.GetDBFromCtx(this.GetCtx())
	qs := db.Model(&db_file.File{}).Where(ghost.Map{
		"user_id": this.UserId,
		"type":    this.Type,
		"path":    this.Path,
		"name":    name,
	})
	if qs.Exist() {
		panic(ghost.NewBusinessError("名称已存在"))
	}

	if result := qs.Update("name", name); result.Error != nil {
		ghost.Error(result.Error)
		panic(ghost.NewSystemError("修改名称失败"))
	}
}

func NewFileFromDbModel(ctx context.Context, dbModel *db_file.File) *File {
	inst := new(File)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}

func NewDir(ctx context.Context, user *bm_account.User, path, dirName string) {
	db := ghost.GetDBFromCtx(ctx)
	if db.Model(&db_file.File{}).Where(ghost.Map{
		"user_id": user.Id,
		"type":    db_file.FILE_TYPE_DIR,
		"path":    path,
		"name":    dirName,
	}).Exist() {
		panic(ghost.NewBusinessError("目录已存在"))
	}
	result := db.Create(&db_file.File{
		UserId: user.Id,
		Type:   db_file.FILE_TYPE_DIR,
		Path:   path,
		Status: db_file.FILE_STATUS_COMPLETE,
		Name:   dirName,
	})
	if result.Error != nil {
		ghost.Error(result.Error)
		panic(ghost.NewSystemError("添加目录失败"))
	}
}
