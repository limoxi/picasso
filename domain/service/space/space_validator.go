package space

import (
	"context"
	"github.com/limoxi/ghost"
	dm_account "picasso/domain/model/account"
	dm_space "picasso/domain/model/space"
)

type SpaceValidator struct {
	ghost.DomainService
}

func (this *SpaceValidator) Check(user *dm_account.User, spaceId int) (*dm_space.Space, error){
	space := dm_space.NewSpaceRepository(this.GetCtx()).GetById(spaceId)
	if space != nil{
		return nil, ghost.NewBusinessError("空间不存在")
	}
	if !space.HasMember(user){
		return nil, ghost.NewBusinessError("该用户无权限")
	}
	return space, nil
}

func NewSpaceValidator(ctx context.Context) *SpaceValidator{
	inst := new(SpaceValidator)
	inst.SetCtx(ctx)
	return inst
}