package db

import (
	"context"

	"service-mutasi/store/postgres_store/sqlc"
)

type IStore interface {
	sqlc.Querier

	TabungTx(ctx context.Context, arg TabungTxParams) (TabungTxResult, error)
	TarikTx(ctx context.Context, arg TarikTxParams) (TarikTxResult, error)
}

type TabungTxParams struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}
type TabungTxResult struct {
	Akun   sqlc.Akun   `json:"akun"`
	Mutasi sqlc.Mutasi `json:"mutasi"`
}

type TarikTxParams struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}
type TarikTxResult struct {
	Akun   sqlc.Akun   `json:"akun"`
	Mutasi sqlc.Mutasi `json:"mutasi"`
}
