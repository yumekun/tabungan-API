package service

import (
	"context"

	"service-akun/db/sqlc"
	"service-akun/dto"
)

func (service *Service) Mutasi(ctx context.Context, request dto.MutasiRequest) (entries []sqlc.Mutasi, err error) {
	
	_, err = service.store.GetAkun(ctx, request.NoRekening)
	if err != nil {
		return nil, err
	}

	entries, err = service.store.GetMutasi(ctx, request.NoRekening)
	if err != nil {
		return nil, err
	}

	return entries, nil
}