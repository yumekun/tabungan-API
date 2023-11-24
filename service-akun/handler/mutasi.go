package handler

import (
	"context"
	"database/sql"

	"service-akun/db/sqlc"
	"service-akun/dto"
	"service-akun/model"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Mutasi(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening", "")

	request := dto.MutasiRequest{
		NoRekening: noRekening,
	}

	// call service layer
	entries, err := handler.service.Mutasi(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.MutasiErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.MutasiErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.MutasiSuccessResponse{
		Mutasi: func(entries []sqlc.Mutasi) []model.Statement {
			mutasi := []model.Statement{}

			// for _, mutasi := range entries {
			// 	mutasi = append(mutasi, model.Statement{
			// 		KodeTransaksi: mutasi.KodeTransaksi,
			// 		Nominal:       mutasi.Nominal,
			// 		Waktu:         mutasi.Waktu,
			// 	})
			// }

			return mutasi
		}(entries),
	}

	return c.JSON(response)
}