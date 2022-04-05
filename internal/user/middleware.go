package user

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) checkGetBalanceInput(c *fiber.Ctx) error {
	userInput := struct {
		ID int `json:"user_id"`
	}{}

	if err := c.BodyParser(&userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "parsing data from request body failed with error: " + err.Error(),
		})
	}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	user, err := h.service.GetUser(customContext, userInput.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": `getting user with that "user_id" from db failed with error: ` + err.Error(),
		})
	}
	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": `there is no user with that "user_id"`,
		})
	}

	c.Locals("user", user)
	return c.Next()
}

func (h *Handler) checkP2PInput(c *fiber.Ctx) error {
	p2pInput := struct {
		FromUserID int `json:"from_user_id"`
		ToUserID   int `json:"to_user_id"`
		Amount     int `json:"amount"`
	}{}

	if err := c.BodyParser(&p2pInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "parsing data from request body failed with error: " + err.Error(),
		})
	}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	fromUser, err := h.service.GetUser(customContext, p2pInput.FromUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": `getting user with that "from_user_id" from db failed with error: ` + err.Error(),
		})
	}
	if fromUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": `there is no user with that "from_user_id"`,
		})
	}

	if fromUser.Balance < p2pInput.Amount {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "not enough balance to make transfer",
		})
	}

	toUser, err := h.service.GetUser(customContext, p2pInput.ToUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": `getting user with that "to_user_id" from db failed with error: ` + err.Error(),
		})
	}
	if toUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": `there is no user with that "to_user_id"`,
		})
	}

	if p2pInput.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"massage": "transfer amount must be positive",
		})
	}

	c.Locals("fromUserID", fromUser.ID)
	c.Locals("toUserID", toUser.ID)
	c.Locals("transferAmount", p2pInput.Amount)
	return c.Next()
}

func (h *Handler) checkBalanceOperationInput(c *fiber.Ctx) error {
	balanceOperationInput := struct {
		UserID int    `json:"user_id"`
		Amount int    `json:"amount"`
		Type   string `json:"type"`
	}{}

	if err := c.BodyParser(&balanceOperationInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "parsing data from request body failed with error: " + err.Error(),
		})
	}

	if balanceOperationInput.Type != "add" && balanceOperationInput.Type != "subtract" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid balance operation type",
		})
	}

	customContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	user, err := h.service.GetUser(customContext, balanceOperationInput.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": `getting user with that "user_id" from db failed with error: ` + err.Error(),
		})
	}
	if user == nil {
		if balanceOperationInput.Type == "subtract" {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": `there is no user with that "user_id"`,
			})
		}
		user = &User{
			ID: balanceOperationInput.UserID,
		}
		err := h.service.CreateUser(customContext, user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"massage": `creating user with that "user_id" in db failed with error: ` + err.Error(),
			})
		}
	}
	c.Locals("userID", user.ID)

	if balanceOperationInput.Amount <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "balance operation amount must be positive",
		})
	}

	if balanceOperationInput.Type == "subtraction" {
		if user.Balance < balanceOperationInput.Amount {
			return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"message": "not enough balance to make operation",
			})
		}
	}

	c.Locals("user", user)
	c.Locals("operationAmount", balanceOperationInput.Amount)
	c.Locals("operationType", balanceOperationInput.Type)
	return c.Next()
}
