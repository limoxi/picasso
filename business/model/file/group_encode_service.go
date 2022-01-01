package file

import (
	"context"
	"github.com/limoxi/ghost"
)

type GroupEncodeService struct {
	ghost.DomainService
}

func (this *GroupEncodeService) Encode(group *Group) *EncodedGroup {
	return &EncodedGroup{
		Id:   group.Id,
		Name: group.Name,
	}
}

func (this *GroupEncodeService) EncodeGroupUser(gUser *GroupUser) *EncodedGroupUser {
	return &EncodedGroupUser{
		UserId:    gUser.UserId,
		NickName:  gUser.NickName,
		IsManager: gUser.IsManager,
	}
}

func (this *GroupEncodeService) EncodeMany(groups []*Group) []*EncodedGroup {
	encodedGroups := make([]*EncodedGroup, 0, len(groups))
	for _, group := range groups {
		encodedGroups = append(encodedGroups, this.Encode(group))
	}
	return encodedGroups
}

func (this *GroupEncodeService) EncodeManyUsers(gUsers []*GroupUser) []*EncodedGroupUser {
	encodedGroupUsers := make([]*EncodedGroupUser, 0, len(gUsers))
	for _, gUser := range gUsers {
		encodedGroupUsers = append(encodedGroupUsers, this.EncodeGroupUser(gUser))
	}
	return encodedGroupUsers
}

func NewGroupEncodeService(ctx context.Context) *GroupEncodeService {
	inst := new(GroupEncodeService)
	inst.SetCtx(ctx)
	return inst
}
