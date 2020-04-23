package account

import (
	"context"
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
)

type LoginService struct{
	ghost.DomainObject
}

type RegisterParams struct{
	Phone string
	RawPassword string
}

func (this *LoginService) Login(phone, password string) *dm_account.User{
	user := dm_account.NewUserRepository(this.GetCtx()).GetByPhone(phone)
	if user == nil{
		panic(ghost.NewBusinessError("用户不存在"))
	}
	if !checkPassword(password, user.Password){
		panic(ghost.NewBusinessError("密码错误"))
	}
	return user
}

// Register 注册
func (this *LoginService) Register(params RegisterParams) *dm_account.User {
	if !this.checkUsername(params.Phone) {
		panic(ghost.NewBusinessError("该手机号已注册"))
	}
	encodedPwd := encodePassword(params.RawPassword)
	return dm_account.NewUserFactory(this.GetCtx()).Create(&dm_account.User{
		Phone:  params.Phone,
		Password:  encodedPwd,
	})
}

func (this *LoginService) checkUsername(phone string) bool{
	return !dm_account.NewUserRepository(this.GetCtx()).UserExisted(phone)
}

func NewLoginService(ctx context.Context) *LoginService{
	inst := new(LoginService)
	inst.SetCtx(ctx)
	return inst
}