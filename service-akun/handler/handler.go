package handler

import (
	"service-akun/service"
)

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		service: service,
	}
}