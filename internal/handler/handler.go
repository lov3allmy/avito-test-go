package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lov3allmy/avito-test-go/internal/domain"
)

type Handler struct {
	service domain.Service
}

func NewHandler(service domain.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) MakeP2PTransfer(c *fiber.Ctx) error {
	p2pInput := c.Locals("p2pInput").(domain.P2PInput)

	err := h.service.MakeP2PTransfer(p2pInput)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "making transfer failed with error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "transfer completed",
	})
}

func (h *Handler) GetBalanceByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.User)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"balance": user.Balance,
	})
}

func (h *Handler) MakeBalanceOperationByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(*domain.User)
	amount := c.Locals("operationAmount").(int)
	operationType := c.Locals("operationType").(string)

	switch operationType {
	case "add":
		user.Balance += amount
		break
	case "subtract":
		user.Balance -= amount
		break
	}

	if err := h.service.UpdateUser(user.ID, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "making operation failed with error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "operation completed",
	})
}
