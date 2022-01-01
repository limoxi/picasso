package file

import (
	"github.com/limoxi/ghost"
	"time"
)

type Group struct {
	ghost.BaseDBModel

	Name          string `gorm:"size:128"`
	UserId        int
	Code          string `gorm:"size:64;default('')"`
	CodeExpiredAt time.Time
}

func (Group) TableName() string {
	return "file_group"
}

type GroupUser struct {
	ghost.BaseDBModel

	GroupId   int
	UserId    int
	IsManager bool
	NickName  string `gorm:"size:128;default('')"`
}

func (GroupUser) TableName() string {
	return "file_group_user"
}

func init() {
	ghost.RegisterDBModel(&Group{})
	ghost.RegisterDBModel(&GroupUser{})
}
