package dto

type SaldoRequest struct {
	NoRekening string `json:"no_rekening"`
}

type SaldoErrorResponse struct {
	Remark string `json:"remark"`
}

type SaldoSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}