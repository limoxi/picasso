package user

import (
	"github.com/limoxi/ghost"
)

// User 用户信息
type User struct {
	ghost.BaseModel
	Phone string `gorm:"size:20;unique"`
	Password string `gorm:"size:128"`
}

func (User) TableName() string{
	return "user_user"
}

func init(){
	ghost.RegisterDBModel(&User{})
}