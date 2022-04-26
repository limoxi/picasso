package account

import (
	"context"
	"github.com/limoxi/ghost"
)

type UserEncodeService struct {
	ghost.DomainService
}

func (this *UserEncodeService) Encode(user *User) *EncodedUser {
	return &EncodedUser{
		Id:       user.Id,
		Avatar:   user.Avatar,
		Nickname: user.Nickname,
		Phone:    user.Phone,
	}
}

func (this *UserEncodeService) EncodeMany(users []*User) []*EncodedUser {
	encodedUsers := make([]*EncodedUser, 0, len(users))
	for _, user := range users {
		encodedUsers = append(encodedUsers, this.Encode(user))
	}
	return encodedUsers
}

func NewUserEncodeService(ctx context.Context) *UserEncodeService {
	inst := new(UserEncodeService)
	inst.SetCtx(ctx)
	return inst
}
