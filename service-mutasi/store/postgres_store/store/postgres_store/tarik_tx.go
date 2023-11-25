package postgres_store

import (
	"context"
	"fmt"

	"service-mutasi/store/postgres_store/sqlc"
	db "service-mutasi/store/postgres_store/store"
)

func (store *PostgresStore) TarikTx(ctx context.Context, arg db.TarikTxParams) (db.TarikTxResult, error) {
	var result db.TarikTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		result.Mutasi, err = q.CreateMutasi(ctx, sqlc.CreateMutasiParams{
			KodeTransaksi: "D",
			Nominal:       arg.Nominal,
			NoRekening:    arg.NoRekening,
		})
		if err != nil {
			fmt.Printf("failed to execute 'CreateMutasi' query: %s", err)

			return err
		}

		akun, err := q.GetAkun(ctx, arg.NoRekening)
		if err != nil {
			fmt.Printf("failed to execute 'GetAkun' query: %s", err)

			return err
		}

		if akun.Saldo < arg.Nominal {
			fmt.Printf("insufficient balance: %s", err)

			return err
		}

		result.Akun, err = q.UpdateSaldo(ctx, sqlc.UpdateSaldoParams{
			NoRekening: arg.NoRekening,
			Saldo:      akun.Saldo - arg.Nominal,
		})
		if err != nil {
			fmt.Printf("failed to execute 'UpdateSaldo' query: %s", err)

			return err
		}

		return err
	})

	return result, err
}
