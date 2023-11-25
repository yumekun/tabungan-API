// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"context"
)

type Querier interface {
	CreateAkun(ctx context.Context, arg CreateAkunParams) (Akun, error)
	CreateNasabah(ctx context.Context, arg CreateNasabahParams) (Nasabah, error)
	GetAkun(ctx context.Context, noRekening string) (Akun, error)
	GetMutasi(ctx context.Context, noRekening string) ([]Mutasi, error)
}

var _ Querier = (*Queries)(nil)