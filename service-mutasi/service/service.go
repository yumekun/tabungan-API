package service

import (
	"context"

	"service-mutasi/dto"
	db "service-mutasi/store/postgres_store/store"
)

type IService interface {
	Tarik(ctx context.Context, request dto.TarikRequest) (saldo int64, err error)
	Tabung(ctx context.Context, request dto.TabungRequest) (saldo int64, err error)
}

type store struct {
	postgres db.IStore
}

func newStore(
	postgresStore db.IStore,
) store {
	return store{
		postgres: postgresStore,
	}
}

type Service struct {
	store store
}

func NewService(
	postgresStore db.IStore,
) IService {
	store := newStore(postgresStore)

	return &Service{
		store: store,
	}
}