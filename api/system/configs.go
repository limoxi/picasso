package system

import (
	"github.com/limoxi/ghost"
	bm_system "picasso/business/model/system"
)

type Configs struct {
	ghost.ApiTemplate
}

func (this *Configs) Resource() string {
	return "system.configs"
}

func (this *Configs) Get() ghost.Response {
	ctx := this.GetCtx()
	configs := bm_system.NewSystemConfigRepository(ctx).GetAll()
	return ghost.NewJsonResponse(ghost.Map{
		"configs": bm_system.NewSystemConfigEncodeService(ctx).EncodeMany(configs),
	})
}

func init() {
	ghost.RegisterApi(&Configs{})
}
