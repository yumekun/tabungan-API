package handler

import (
	"context"

	"service-akun/dto"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

func (handler *Handler) Daftar(c *fiber.Ctx) error {
	var request dto.DaftarRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.DaftarErrorResponse{
			Remark: "failed to parse request body",
		})
	}


	noRekening, err := handler.service.Daftar(context.Background(), request)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				return c.Status(fiber.StatusBadRequest).JSON(&dto.DaftarErrorResponse{
					Remark: "nik atau nomor hp sudah digunakan",
				})
			}
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.DaftarErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.DaftarSuccessResponse{
		NoRekening: noRekening,
	}

	return c.JSON(response)
}