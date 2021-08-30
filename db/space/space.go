package space

import (
	"github.com/limoxi/ghost"
	"time"
)

type Space struct {
	ghost.BaseDBModel

	Name string `gorm:"size:128"`
	UserId int
	Code string `gorm:"size:64;default('')"`
	CodeExpiredAt time.Time
}
func (Space) TableName() string{
	return "space_space"
}

type SpaceMember struct {
	ghost.BaseDBModel

	SpaceId int
	UserId int
	IsManager bool
	NickName string `gorm:"size:128;default('')"`
}
func (SpaceMember) TableName() string{
	return "space_member"
}

func init(){
	ghost.RegisterDBModel(&Space{})
	ghost.RegisterDBModel(&SpaceMember{})
}