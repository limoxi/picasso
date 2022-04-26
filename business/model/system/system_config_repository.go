package system

import (
	"context"
	"github.com/limoxi/ghost"
	db_system "picasso/db/system"
)

type SystemConfigRepository struct {
	ghost.DomainService
}

func (this *SystemConfigRepository) GetByFilters(filters ghost.Map) []*SystemConfig {
	if filters == nil {
		filters = make(ghost.Map)
	}
	var dbModels []*db_system.Config
	configs := make([]*SystemConfig, 0)
	ctx := this.GetCtx()
	result := ghost.GetDBFromCtx(ctx).Where(filters).Find(&dbModels)
	if err := result.Error; err != nil {
		ghost.Error(err.Error())
		return configs
	}
	for _, dbModel := range dbModels {
		configs = append(configs, NewSystemConfigFromDbModel(ctx, dbModel))
	}
	return configs
}

func (this *SystemConfigRepository) GetAll() []*SystemConfig {
	return this.GetByFilters(nil)
}

func NewSystemConfigRepository(ctx context.Context) *SystemConfigRepository {
	inst := new(SystemConfigRepository)
	inst.SetCtx(ctx)
	return inst
}
