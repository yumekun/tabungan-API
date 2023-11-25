package service

import (
	"context"

	"service-mutasi/dto"
	db "service-mutasi/store/postgres_store/store"
)

func (service *Service) Tabung(ctx context.Context, request dto.TabungRequest) (saldo int64, err error) {

	result, err := service.store.postgres.TabungTx(ctx, db.TabungTxParams{
		Nominal:    request.Nominal,
		NoRekening: request.NoRekening,
	})
	if err != nil {
		return -1, err
	}

	return result.Akun.Saldo, nil
}