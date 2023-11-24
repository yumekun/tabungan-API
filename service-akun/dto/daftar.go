package dto

type DaftarRequest struct {
	Nama string `json:"nama"`
	Nik  string `json:"nik"`
	NoHp string `json:"no_hp"`
}

type DaftarErrorResponse struct {
	Remark string `json:"remark"`
}

type DaftarSuccessResponse struct {
	NoRekening string `json:"no_rekening"`
}