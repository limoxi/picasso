package space

import (
	"context"
	"github.com/limoxi/ghost"
	ghost_util "github.com/limoxi/ghost/utils"
	"math/rand"
	db_space "picasso/db/space"
	bm_account "picasso/business/model/account"
	"strconv"
	"time"
)

type Space struct {
	ghost.DomainModel
	Id int
	Name string
	Code string
	CodeExpiredAt time.Time
}

func (this *Space) checkCode(code string){
	if code != this.Code{
		panic(ghost.NewBusinessError("无效的邀请码"))
	}
	ghost.Info(time.Now(), this.CodeExpiredAt, time.Now().After(this.CodeExpiredAt))
	if time.Now().After(this.CodeExpiredAt){
		panic(ghost.NewBusinessError("邀请码已过期"))
	}
}

func (this *Space) AddMember(member *bm_account.User, code string){
	if this.HasMember(member){
		panic(ghost.NewBusinessError("用户已加入"))
	}

	this.checkCode(code)

	if err := ghost.GetDBFromCtx(this.GetCtx()).Create(&db_space.SpaceMember{
		SpaceId: this.Id,
		UserId: member.Id,
		IsManager: false,
	}).Error; err != nil{
		ghost.Panic(err)
	}
}

func (this *Space) HasMember(member *bm_account.User) bool{
	filters := ghost.Map{
		"space_id": this.Id,
		"user_id": member.Id,
	}
	var count int
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_space.SpaceMember{}).Where(filters).Count(&count)
	if err := result.Error; err != nil{
		panic(err)
	}
	return count > 0
}

// GetMembers 获取成员
func (this *Space) GetMembers() []*SpaceMember{
	ctx := this.GetCtx()
	filters := ghost.Map{
		"space_id": this.Id,
	}
	var dbModels []*db_space.SpaceMember
	result := ghost.GetDBFromCtx(ctx).Model(&db_space.SpaceMember{}).Where(filters).Order("-id").Find(&dbModels)
	if err := result.Error; err != nil{
		panic(err)
	}
	members := make([]*SpaceMember, 0, len(dbModels))
	for _, dbModel := range dbModels{
		members = append(members, NewSpaceMemberFromDbModel(ctx, dbModel))
	}
	return members
}

// GenCode 随机生成4位数字邀请码
func (this *Space) GenCode() string{
	rand.Seed(time.Now().Unix())
	code := strconv.Itoa(rand.Intn(8999) + 1000)
	result := ghost.GetDBFromCtx(this.GetCtx()).Model(&db_space.Space{}).Where(ghost.Map{
		"id": this.Id,
	}).Update(ghost.Map{
		"code": code,
		"code_expired_at": time.Now().Add(time.Hour * 24),
	})
	if err := result.Error; err != nil{
		panic(err)
	}
	return code
}

func NewSpaceFromDbModel(ctx context.Context, dbModel *db_space.Space) *Space{
	inst := new(Space)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}

func NewSpaceForUser(ctx context.Context, user *bm_account.User, name string) *Space{
	dbModel := &db_space.Space{
		Name: name,
		UserId: user.Id,
		CodeExpiredAt: ghost_util.DEFAULT_TIME,
	}
	db := ghost.GetDBFromCtx(ctx)
	result := db.Create(dbModel)
	if err := result.Error; err != nil{
		panic(ghost.NewSystemError(err.Error(), "创建空间失败"))
	}
	result = db.Create(&db_space.SpaceMember{
		SpaceId: dbModel.Id,
		UserId: user.Id,
		IsManager: true,
	})
	if err := result.Error; err != nil{
		panic(ghost.NewSystemError(err.Error(), "增加空间成员失败"))
	}
	return &Space{
		Id: dbModel.Id,
	}
}
