package service

import (
	"context"

	"service-mutasi/dto"
	db "service-mutasi/store/postgres_store/store"
)

func (service *Service) Tarik(ctx context.Context, request dto.TarikRequest) (saldo int64, err error) {

	result, err := service.store.postgres.TarikTx(ctx, db.TarikTxParams{
		Nominal:    request.Nominal,
		NoRekening: request.NoRekening,
	})
	if err != nil {
		return -1, err
	}

	return result.Akun.Saldo, nil
}