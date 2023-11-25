package service

import (
	"context"

	"service-akun/dto"
	"service-akun/store/postgres_store/sqlc"
)

func (service *Service) Mutasi(ctx context.Context, request dto.MutasiRequest) (entries []sqlc.Mutasi, err error) {
	
	_, err = service.store.postgres.GetAkun(ctx, request.NoRekening)
	if err != nil {
		return nil, err
	}

	entries, err = service.store.postgres.GetMutasi(ctx, request.NoRekening)
	if err != nil {
		return nil, err
	}

	return entries, nil
}