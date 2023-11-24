package handler

import (
	"context"
	"database/sql"

	"service-akun/dto"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) Saldo(c *fiber.Ctx) error {
	noRekening := c.Params("no_rekening", "")

	request := dto.SaldoRequest{
		NoRekening: noRekening,
	}

	saldo, err := handler.service.Saldo(context.Background(), request)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(fiber.StatusBadRequest).JSON(&dto.SaldoErrorResponse{
				Remark: "nomor rekening tidak dikenali`",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&dto.SaldoErrorResponse{
			Remark: "internal server error",
		})
	}

	response := &dto.SaldoSuccessResponse{
		Saldo: saldo,
	}

	return c.JSON(response)
}