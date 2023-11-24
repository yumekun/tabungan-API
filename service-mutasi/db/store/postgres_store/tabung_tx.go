package postgres_store

import (
	"context"
	"fmt"

	"service-mutasi/db/sqlc"
	db "service-mutasi/db/store"
)

func (store *PostgresStore) TabungTx(ctx context.Context, arg db.TabungTxParams) (db.TabungTxResult, error) {

	var result db.TabungTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error


		result.Mutasi, err = q.CreateMutasi(ctx, sqlc.CreateMutasiParams{
			KodeTransaksi:       "C",
			Nominal:    arg.Nominal,
			NoRekening: arg.NoRekening,
		})
		if err != nil {
			fmt.Printf("failed to execute 'CreateMutasi' query: %s", err)


			return err
		}

		akun, err := q.GetAkun(ctx, arg.NoRekening)
		if err != nil {
			fmt.Printf("failed to execute 'GetAccount' query: %s", err)


			return err
		}

		result.Akun, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NoRekening: arg.NoRekening,
			Saldo:      akun.Saldo + arg.Nominal,
		})
		if err != nil {
			fmt.Printf("failed to execute 'UpdateSaldo' query: %s", err)

		

			return err
		}

		return err
	})

	return result, err
}