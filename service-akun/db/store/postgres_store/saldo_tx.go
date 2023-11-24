package postgres_store

import (
	"context"
	"fmt"

	"service-akun/db/sqlc"
	db "service-akun/db/store"
)


func (store *PostgresStore) SaldoTx(ctx context.Context, arg db.SaldoTxParams) (db.SaldoTxResult, error) {

	var result db.SaldoTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		result.Akun, err = q.GetAkun(ctx, arg.NoRekening)
		if err != nil {
			fmt.Printf("failed to execute 'GetNasabah' query: %s", err)
			return err
		}

		return err
	})

	return result, err
}