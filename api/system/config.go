package system

import (
	"github.com/limoxi/ghost"
	bm_system "picasso/business/model/system"
)

type Config struct {
	ghost.ApiTemplate

	PostParams *struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
}

func (this *Config) Resource() string {
	return "system.config"
}

func (this *Config) Post() ghost.Response {
	ctx := this.GetCtx()
	params := this.PostParams
	bm_system.SaveConfig(ctx, params.Key, params.Value)
	return ghost.NewJsonResponse(nil)
}

func init() {
	ghost.RegisterApi(&Config{})
}
