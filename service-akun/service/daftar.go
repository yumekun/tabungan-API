package service

import (
	"context"

	"service-akun/dto"
	db "service-akun/store/postgres_store/store"
)

func (service *Service) Daftar(ctx context.Context, request dto.DaftarRequest) (noRekening string, err error) {

	result, err := service.store.postgres.DaftarTx(ctx, db.DaftarTxParams{
		Nama: request.Nama,
		Nik:  request.Nik,
		NoHp: request.NoHp,
	})
	if err != nil {
		return "", err
	}

	return result.Akun.NoRekening, nil
}