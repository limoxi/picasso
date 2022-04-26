package account

import (
	"context"
	"github.com/limoxi/ghost"
	m_account "picasso/business/model/account"
)

type LoginService struct {
	ghost.DomainService
}

type RegisterParams struct {
	Avatar      string
	Nickname    string
	Phone       string
	RawPassword string
}

func (this *LoginService) Login(phone, password string) *m_account.User {
	user := m_account.NewUserRepository(this.GetCtx()).GetByPhone(phone)
	if user == nil {
		panic(ghost.NewBusinessError("用户不存在"))
	}
	if !checkPassword(password, user.Password) {
		panic(ghost.NewBusinessError("密码错误"))
	}
	return user
}

// Register 注册
func (this *LoginService) Register(params RegisterParams) *m_account.User {
	if !this.checkPhone(params.Phone) {
		panic(ghost.NewBusinessError("该手机号已注册"))
	}
	encodedPwd := encodePassword(params.RawPassword)
	return m_account.NewUserFactory(this.GetCtx()).Create(&m_account.User{
		Avatar:   params.Avatar,
		Nickname: params.Nickname,
		Phone:    params.Phone,
		Password: encodedPwd,
	})
}

func (this *LoginService) checkPhone(phone string) bool {
	return !m_account.NewUserRepository(this.GetCtx()).UserExisted(phone)
}

func NewLoginService(ctx context.Context) *LoginService {
	inst := new(LoginService)
	inst.SetCtx(ctx)
	return inst
}
