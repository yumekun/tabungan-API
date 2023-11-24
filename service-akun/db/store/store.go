package db

import (
	"context"

	"service-akun/db/sqlc"
)
type IStore interface {
	sqlc.Querier
	DaftarTx(ctx context.Context, arg DaftarTxParams) (DaftarTxResult, error)
	SaldoTx(ctx context.Context,arg SaldoTxParams) (SaldoTxResult,error)
	MutasiTx(ctx context.Context,arg MutasiTxParams) (MutasiTxResult,error)
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

type SaldoTxParams struct {
	NoRekening string `json:"no_rekening"`
}

type SaldoTxResult struct {
	Akun  sqlc.Akun  `json:"akun"`
}

type MutasiTxParams struct {
	NoRekening string `json:"no_rekening"`
}

type MutasiTxResult struct {
	Mutasi  []sqlc.Mutasi  `json:"mutasi"`
}