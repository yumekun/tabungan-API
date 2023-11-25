package dto

type TabungRequest struct {
	NoRekening string `json:"no_rekening"`
	Nominal    int64  `json:"nominal"`
}

type TabungErrorResponse struct {
	Remark string `json:"remark"`
}

type TabungSuccessResponse struct {
	Saldo int64 `json:"saldo"`
}