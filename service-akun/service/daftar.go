package service

import (
	"context"

	db "service-akun/db/store"
	"service-akun/dto"
)

func (service *Service) Daftar(ctx context.Context, request dto.DaftarRequest) (noRekening string, err error) {

	result, err := service.store.DaftarTx(ctx, db.DaftarTxParams{
		Nama: request.Nama,
		Nik:  request.Nik,
		NoHp: request.NoHp,
	})
	if err != nil {
		return "", err
	}

	return result.Akun.NoRekening, nil
}