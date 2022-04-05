package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service Service
}

func NewHandler(router fiber.Router, service Service) {
	handler := &Handler{
		service: service,
	}
	balance := router.Group("/balance")
	balance.Get("", handler.checkGetBalanceInput, handler.getBalanceByUserID)
	balance.Post("", handler.checkBalanceOperationInput, handler.makeBalanceOperationByUserID)
	router.Post("/p2p", handler.checkP2PInput, handler.makeP2PTransfer)
}

func (h *Handler) makeP2PTransfer(c *fiber.Ctx) error {
	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	fromUserID := c.Locals("fromUserID").(int)
	toUserID := c.Locals("toUserID").(int)
	amount := c.Locals("transferAmount").(int)

	err := h.service.makeP2PTransfer(customContext, fromUserID, toUserID, amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "making transfer failed with error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "transfer completed",
	})
}

func (h *Handler) getBalanceByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(*User)

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"balance": user.Balance,
	})
}

func (h *Handler) makeBalanceOperationByUserID(c *fiber.Ctx) error {
	user := c.Locals("user").(*User)
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

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := h.service.UpdateUser(customContext, user.ID, user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "making operation failed with error: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"message": "operation completed",
	})
}
