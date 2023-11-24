package dto

import (
	"service-akun/model"
)

type MutasiRequest struct {
	NoRekening string `json:"no_rekening"`
}

type MutasiErrorResponse struct {
	Remark string `json:"remark"`
}

type MutasiSuccessResponse struct {
	Mutasi []model.Statement `json:"mutasi"`
}