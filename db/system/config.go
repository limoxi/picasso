package user

import (
	"github.com/limoxi/ghost"
)

// Config 配置信息
type Config struct {
	ghost.BaseDBModel
	K string `gorm:"size:32"`
	V string `gorm:"size:256"`
}

func (Config) TableName() string {
	return "system_config"
}

func init() {
	ghost.RegisterDBModel(&Config{})
}
