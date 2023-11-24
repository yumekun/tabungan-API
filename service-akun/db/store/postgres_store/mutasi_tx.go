package postgres_store

import (
	"context"
	"fmt"

	"service-akun/db/sqlc"
	db "service-akun/db/store"
)


func (store *PostgresStore) MutasiTx(ctx context.Context, arg db.MutasiTxParams) (db.MutasiTxResult, error) {

	var result db.MutasiTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		result.Mutasi, err = q.GetMutasi(ctx, arg.NoRekening)
		if err != nil {
			fmt.Printf("failed to execute 'GetMutasi' query: %s", err)
			return err
		}

		return err
	})

	return result, err
}