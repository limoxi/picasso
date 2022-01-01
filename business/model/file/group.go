package file

import (
	"context"
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"math/rand"
	"picasso/business/model/account"
	db_file "picasso/db/file"
	"strconv"
	"time"
)

type Group struct {
	ghost.DomainModel
	Id            int
	Name          string
	Code          string
	CodeExpiredAt time.Time
}

func (this *Group) checkCode(code string) {
	if code != this.Code {
		panic(ghost.NewBusinessError("无效的邀请码"))
	}
	ghost.Info(time.Now(), this.CodeExpiredAt, time.Now().After(this.CodeExpiredAt))
	if time.Now().After(this.CodeExpiredAt) {
		panic(ghost.NewBusinessError("邀请码已过期"))
	}
}

func (this *Group) AddUser(user *account.User, code string) {
	if this.HasUser(user) {
		panic(ghost.NewBusinessError("用户已加入"))
	}

	this.checkCode(code)

	if err := ghost.GetDBFromCtx(this.GetCtx()).Create(&db_file.GroupUser{
		GroupId:   this.Id,
		UserId:    user.Id,
		IsManager: false,
	}).Error; err != nil {
		panic(err)
	}
}

func (this *Group) HasUser(user *account.User) bool {
	filters := ghost.Map{
		"group_id": this.Id,
		"user_id":  user.Id,
	}
	var count int64
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_file.GroupUser{}).Where(filters).Count(&count)
	if err := result.Error; err != nil {
		panic(err)
	}
	return count > 0
}

func (this *Group) GetUsers() []*GroupUser {
	ctx := this.GetCtx()
	filters := ghost.Map{
		"group_id": this.Id,
	}
	var dbModels []*db_file.GroupUser
	result := ghost.GetDBFromCtx(ctx).Model(&db_file.GroupUser{}).Where(filters).Order("-id").Find(&dbModels)
	if err := result.Error; err != nil {
		panic(err)
	}
	gUsers := make([]*GroupUser, 0, len(dbModels))
	for _, dbModel := range dbModels {
		gUsers = append(gUsers, NewGroupUserFromDbModel(ctx, dbModel))
	}
	return gUsers
}

// GenCode 随机生成4位数字邀请码
func (this *Group) GenCode() string {
	rand.Seed(time.Now().Unix())
	code := strconv.Itoa(rand.Intn(8999) + 1000)
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_file.Group{}).Where(ghost.Map{
		"id": this.Id,
	}).Updates(ghost.Map{
		"code":            code,
		"code_expired_at": time.Now().Add(time.Hour * 24),
	})
	if err := result.Error; err != nil {
		panic(err)
	}
	return code
}

func NewGroupFromDbModel(ctx context.Context, dbModel *db_file.Group) *Group {
	inst := new(Group)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}

func NewGroupForUser(ctx context.Context, user *account.User, name string) *Group {
	dbModel := &db_file.Group{
		Name:          name,
		UserId:        user.Id,
		CodeExpiredAt: ghost_util.DEFAULT_TIME,
	}
	db := ghost.GetDBFromCtx(ctx)
	result := db.Create(dbModel)
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError(err.Error(), "创建失败"))
	}
	result = db.Create(&db_file.GroupUser{
		GroupId:   dbModel.Id,
		UserId:    user.Id,
		IsManager: true,
	})
	if err := result.Error; err != nil {
		panic(ghost.NewSystemError(err.Error(), "增加成员失败"))
	}
	return &Group{
		Id: dbModel.Id,
	}
}
