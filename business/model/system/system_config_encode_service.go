package system

import (
	"context"
	"github.com/limoxi/ghost"
)

type SystemConfigEncodeService struct {
	ghost.DomainService
}

func (this *SystemConfigEncodeService) Encode(config *SystemConfig) *EncodedConfig {
	return &EncodedConfig{
		K: config.K,
		V: config.V,
	}
}

func (this *SystemConfigEncodeService) EncodeMany(configs []*SystemConfig) []*EncodedConfig {
	encodedRecords := make([]*EncodedConfig, 0, len(configs))
	for _, config := range configs {
		encodedRecords = append(encodedRecords, this.Encode(config))
	}
	return encodedRecords
}

func NewSystemConfigEncodeService(ctx context.Context) *SystemConfigEncodeService {
	inst := new(SystemConfigEncodeService)
	inst.SetCtx(ctx)
	return inst
}
