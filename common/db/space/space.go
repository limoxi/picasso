package space

import "github.com/limoxi/ghost"

type Space struct {
	ghost.BaseModel

	Name string `gorm:"size:128"`
	UserId int
}

func (Space) TableName() string{
	return "space_space"
}

func init(){
	ghost.RegisterDBModel(&Space{})
}