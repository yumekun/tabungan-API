package dto

type TarikRequest struct {
	NoRekening string `json:"no_rekening"`
	Nominal    int64  `json:"nominal"`
}

type TarikErrorResponse struct {
	Remark string `json:"remark"`
}

type TarikSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}