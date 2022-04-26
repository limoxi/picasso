package user

import (
	"github.com/limoxi/ghost"
)

// User 用户信息
type User struct {
	ghost.BaseDBModel
	Avatar   string `gorm:"size:256"`
	Nickname string `gorm:"size:24"`
	Phone    string `gorm:"size:20;unique"`
	Password string `gorm:"size:64"`
}

func (User) TableName() string {
	return "auth_user"
}

func init() {
	ghost.RegisterDBModel(&User{})
}
