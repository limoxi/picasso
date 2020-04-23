package account

import (
	"context"
	"github.com/limoxi/ghost"
	ghost_utils "github.com/limoxi/ghost/utils"
)

type UserEncodeService struct {
	ghost.DomainObject
}

func (this *UserEncodeService) Encode(user *User) *EncodedUser{
	return &EncodedUser{
		Id:   user.Id,
		Phone: user.Phone,
		CreatedAt: ghost_utils.FormatDatetime(user.CreatedAt),
	}
}

func (this *UserEncodeService) EncodeMany(users []*User) []*EncodedUser{
	encodedUsers := make([]*EncodedUser, 0, len(users))
	for _, user := range users{
		encodedUsers = append(encodedUsers, this.Encode(user))
	}
	return encodedUsers
}

func NewUserEncodeService(ctx context.Context) *UserEncodeService{
	inst := new(UserEncodeService)
	inst.SetCtx(ctx)
	return inst
}