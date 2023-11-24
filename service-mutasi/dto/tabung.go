package dto

type TabungRequest struct {
	Nominal    int64  `json:"nominal"`
	NoRekening string `json:"no_rekening"`
}

type TabungErrorResponse struct {
	Remark string `json:"remark"`
}

type TabungSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}