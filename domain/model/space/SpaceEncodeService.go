package space

import (
	"context"
	"github.com/limoxi/ghost"
)

type SpaceEncodeService struct {
	ghost.DomainObject
}

func (this *SpaceEncodeService) Encode(space *Space) *EncodedSpace{
	return &EncodedSpace{
		Id: space.Id,
		Name: space.Name,
	}
}

func (this *SpaceEncodeService) EncodeMember(member *SpaceMember) *EncodedSpaceMember{
	return &EncodedSpaceMember{
		UserId: member.UserId,
		NickName: member.NickName,
		IsManager: member.IsManager,
	}
}

func (this *SpaceEncodeService) EncodeMany(spaces []*Space) []*EncodedSpace{
	encodedSpaces := make([]*EncodedSpace, 0, len(spaces))
	for _, space := range spaces{
		encodedSpaces = append(encodedSpaces, this.Encode(space))
	}
	return encodedSpaces
}

func (this *SpaceEncodeService) EncodeManyMembers(members []*SpaceMember) []*EncodedSpaceMember{
	encodedSpaceMembers := make([]*EncodedSpaceMember, 0, len(members))
	for _, member := range members{
		encodedSpaceMembers = append(encodedSpaceMembers, this.EncodeMember(member))
	}
	return encodedSpaceMembers
}

func NewSpaceEncodeService(ctx context.Context) *SpaceEncodeService{
	inst := new(SpaceEncodeService)
	inst.SetCtx(ctx)
	return inst
}