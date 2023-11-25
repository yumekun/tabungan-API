package handler

import (
	"context"
	"database/sql"

	"service-akun/dto"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Tarik(c *fiber.Ctx) error {
	var request dto.TarikRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&dto.TarikErrorResponse{
			Remark: "failed to parse request body",
		})
	}

	saldo, err := handler.service.Tarik(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.SaldoErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.TarikErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.TarikSuccessResponse{
		Saldo: saldo,
	}

	return c.JSON(response)
}