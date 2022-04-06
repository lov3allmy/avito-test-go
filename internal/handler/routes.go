package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Router(api fiber.Router, handler *Handler) {
	api.Get("/balance", handler.CheckGetBalanceInput, handler.GetBalanceByUserID)
	api.Post("/balance", handler.CheckBalanceOperationInput, handler.MakeBalanceOperationByUserID)
	api.Post("/p2p", handler.CheckP2PInput, handler.MakeP2PTransfer)
}
