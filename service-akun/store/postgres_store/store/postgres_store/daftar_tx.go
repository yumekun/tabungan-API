package postgres_store

import (
	"context"
	"fmt"

	"service-akun/store/postgres_store/sqlc"
	db "service-akun/store/postgres_store/store"
	"service-akun/util/random"
)


func (store *PostgresStore) DaftarTx(ctx context.Context, arg db.DaftarTxParams) (db.DaftarTxResult, error) {

	var result db.DaftarTxResult

	err := store.execTx(ctx, func(q *sqlc.Queries) error {
		var err error

		result.Nasabah, err = q.CreateNasabah(ctx, sqlc.CreateNasabahParams{
			Nama: arg.Nama,
			Nik:  arg.Nik,
			NoHp: arg.NoHp,
		})
		if err != nil {
			fmt.Printf("failed to execute 'CreateNasabah' query: %s", err)
			return err
		}

		result.Akun, err = q.CreateAkun(ctx, sqlc.CreateAkunParams{
			NasabahID: result.Nasabah.NasabahID,
			NoRekening: random.GenerateNumericString(16),
		})
		if err != nil {
			fmt.Printf("failed to execute 'CreateAkun' query: %s", err)


			return err
		}

		return err
	})

	return result, err
}