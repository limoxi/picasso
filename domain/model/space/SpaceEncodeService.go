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

func (this *SpaceEncodeService) EncodeMany(spaces []*Space) []*EncodedSpace{
	encodedSpaces := make([]*EncodedSpace, 0, len(spaces))
	for _, space := range spaces{
		encodedSpaces = append(encodedSpaces, this.Encode(space))
	}
	return encodedSpaces
}

func NewSpaceEncodeService(ctx context.Context) *SpaceEncodeService{
	inst := new(SpaceEncodeService)
	inst.SetCtx(ctx)
	return inst
}