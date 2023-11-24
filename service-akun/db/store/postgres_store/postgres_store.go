package postgres_store

import (
	"context"
	"database/sql"
	"fmt"

	"service-akun/db/sqlc"
	db "service-akun/db/store"
	"service-akun/util/errs"
)

type PostgresStore struct {
	*sqlc.Queries
	db     *sql.DB
}

func NewPostgresStore( db *sql.DB) db.IStore {
	return &PostgresStore{
		db:      db,
		Queries: sqlc.New(db),
	}
}

func (store *PostgresStore) execTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	const op errs.Op = "store/PostgresStore.execTx"

	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
	fmt.Printf("failed to begin tx: %s", err)


		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
 			fmt.Printf("failed to rollback tx: %s", err)
			return err
		}
		return err
	}

	return tx.Commit()
}