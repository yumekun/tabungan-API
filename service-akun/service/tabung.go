package service

import (
	"context"

	"service-akun/dto"
)

func (service *Service) Tabung(ctx context.Context, request dto.TabungRequest) (saldo int64, err error) {

	akun, err := service.store.postgres.GetAkun(ctx, request.NoRekening)
	if err != nil {
		return -1, err
	}

	err = service.store.redis.AddToStream(ctx, service.config.RedisStreamRequest, map[string]interface{}{
		"req_type":    "tabung",
		"no_rekening": request.NoRekening,
		"nominal":     request.Nominal,
	})
	if err != nil {
		return -1, err
	}

	return akun.Saldo, nil
}