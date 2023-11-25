package service

import (
	"context"

	"service-akun/dto"
	"service-akun/store/postgres_store/sqlc"
	db "service-akun/store/postgres_store/store"
	"service-akun/store/redis_store"
)

type IService interface {
	Daftar(ctx context.Context, request dto.DaftarRequest) (noRekening string, err error)
	Mutasi(ctx context.Context, request dto.MutasiRequest) (entries []sqlc.Mutasi, err error)
	Saldo(ctx context.Context, request dto.SaldoRequest) (saldo int64, err error)
}

type store struct {
	postgres db.IStore
	redis *redis_store.RedisStore
}

func newStore(
	postgresStore db.IStore, 
	redisStore *redis_store.RedisStore,
) store{
	return store{
		postgres: postgresStore,
		redis: redisStore,
	}
}

type Service struct {
	store  store
}

func NewService(
	postgresStore db.IStore,
	redisStore *redis_store.RedisStore,
) IService {
	store := newStore(postgresStore, redisStore)

	return &Service{
		store:  store,
	}
}