package handler

import (
	"context"
	"database/sql"

	"service-akun/dto"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Tabung(c *fiber.Ctx) error {
	var request dto.TabungRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.TabungErrorResponse{
			Remark: "failed to parse request body",
		})
	}

	saldo, err := handler.service.Tabung(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.SaldoErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.TabungErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.TabungSuccessResponse{
		Saldo: saldo,
	}

	return c.JSON(response)
}