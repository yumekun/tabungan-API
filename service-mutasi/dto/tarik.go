package dto

type TarikRequest struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type TarikErrorResponse struct {
	Remark string `json:"remark"`
}

type TarikSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}