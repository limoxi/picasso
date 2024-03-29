package user

import (
	"github.com/limoxi/ghost"
)

// User 用户信息
type User struct {
	ghost.BaseDBModel
	Phone string `gorm:"size:20;unique"`
	Password string `gorm:"size:128"`
}

func (User) TableName() string{
	return "auth_user"
}

func init(){
	ghost.RegisterDBModel(&User{})
}