package db

import (
	"context"

	"service-akun/store/postgres_store/sqlc"
)
type IStore interface {
	sqlc.Querier
	DaftarTx(ctx context.Context, arg DaftarTxParams) (DaftarTxResult, error)
}
type DaftarTxParams struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type DaftarTxResult struct {
	Akun  sqlc.Akun  `json:"akun"`
	Nasabah sqlc.Nasabah `json:"nasabah"`
}

