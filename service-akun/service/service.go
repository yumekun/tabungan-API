package service

import (
	"context"

	"service-akun/dto"
	"service-akun/store/postgres_store/sqlc"
	db "service-akun/store/postgres_store/store"
)

type IService interface {
	Daftar(ctx context.Context, request dto.DaftarRequest) (noRekening string, err error)
	Mutasi(ctx context.Context, request dto.MutasiRequest) (entries []sqlc.Mutasi, err error)
	Saldo(ctx context.Context, request dto.SaldoRequest) (saldo int64, err error)
}

type Service struct {
	
	store  db.IStore
}

func NewService( store db.IStore) IService {
	return &Service{
		store:  store,
	}
}