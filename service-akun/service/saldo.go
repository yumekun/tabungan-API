package service

import (
	"context"

	"service-akun/dto"
)

func (service *Service) Saldo(ctx context.Context, request dto.SaldoRequest) (saldo int64, err error) {

	akun, err := service.store.postgres.GetAkun(ctx, request.NoRekening)
	if err != nil {
		return -1, err
	}

	return akun.Saldo, nil
}