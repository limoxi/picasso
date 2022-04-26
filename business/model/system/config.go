package system

import (
	"context"
	"github.com/limoxi/ghost"
	db_system "picasso/db/system"
)

type SystemConfig struct {
	ghost.DomainModel

	K string
	V string
}

func NewSystemConfigFromDbModel(ctx context.Context, dbModel *db_system.Config) *SystemConfig {
	inst := new(SystemConfig)
	inst.SetCtx(ctx)
	inst.NewFromDbModel(inst, dbModel)
	return inst
}

func SaveConfig(ctx context.Context, k, v string) {
	
}
