package space

import (
	"context"
	"github.com/limoxi/ghost"
	"math/rand"
	m_space "picasso/common/db/space"
	dm_account "picasso/domain/model/account"
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
	if time.Now().After(this.CodeExpiredAt){
		panic(ghost.NewBusinessError("邀请码已过期"))
	}
}

func (this *Space) AddUser(user *dm_account.User, code string){
	if this.HasUser(user){
		panic(ghost.NewBusinessError("用户已加入"))
	}

	this.checkCode(code)

	if err := ghost.GetDB().Create(&m_space.SpaceHasUser{
		SpaceId: this.Id,
		UserId: user.Id,
		IsManager: false,
	}).Error; err != nil{
		ghost.Panic(err)
	}
}

func (this *Space) HasUser(user *dm_account.User) bool{
	filters := ghost.Map{
		"space_id": this.Id,
		"user_id": user.Id,
	}
	var count int
	result := ghost.GetDB().Model(&m_space.SpaceHasUser{}).Where(filters).Count(&count)
	if err := result.Error; err != nil{
		panic(err)
	}
	return count > 0
}

// GenCode 随机生成4位数字邀请码
func (this *Space) GenCode() string{
	rand.Seed(time.Now().Unix())
	code := strconv.Itoa(rand.Intn(8999) + 1000)
	result := ghost.GetDB().Model(&m_space.Space{}).Where(ghost.Map{
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

func NewSpaceFromDbModel(ctx context.Context, dbModel *m_space.Space) *Space{
	inst := new(Space)
	inst.SetCtx(ctx)
	inst.Id = dbModel.Id
	inst.Name = dbModel.Name
	return inst
}

func NewSpaceForUser(ctx context.Context, user *dm_account.User, name string) *Space{
	dbModel := &m_space.Space{
		Name: name,
		UserId: user.Id,
	}
	result := ghost.GetDB().Create(dbModel)
	if err := result.Error; err != nil{
		panic(ghost.NewSystemError(err.Error(), "创建空间失败"))
	}
	return &Space{
		Id: dbModel.Id,
	}
}
