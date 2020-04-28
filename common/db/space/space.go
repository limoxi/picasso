package space

import (
	"github.com/limoxi/ghost"
	"time"
)

type Space struct {
	ghost.BaseModel

	Name string `gorm:"size:128"`
	UserId int
	Code string `gorm:"size:64;default('')"`
	CodeExpiredAt time.Time `gorm:"null"`
}
func (Space) TableName() string{
	return "space_space"
}

type SpaceHasUser struct {
	ghost.BaseModel

	SpaceId int
	UserId int
	IsManager bool
}
func (SpaceHasUser) TableName() string{
	return "space_has_user"
}

func init(){
	ghost.RegisterDBModel(&Space{})
	ghost.RegisterDBModel(&SpaceHasUser{})
}