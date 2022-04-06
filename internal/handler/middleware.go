package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lov3allmy/avito-test-go/internal/domain"
)

func (h *Handler) CheckGetBalanceInput(c *fiber.Ctx) error {
	getBalanceInput := domain.GetBalanceInput{}

	if err := c.BodyParser(&getBalanceInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "parsing data from request body failed with error: " + err.Error(),
		})
	}

	if err := ValidateGetBalanceInput(getBalanceInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid request body",
			"errors":  err,
		})
	}

	user, err := h.service.GetUser(getBalanceInput.ID)
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

func (h *Handler) CheckP2PInput(c *fiber.Ctx) error {
	p2pInput := domain.P2PInput{}

	if err := c.BodyParser(&p2pInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "parsing data from request body failed with error: " + err.Error(),
		})
	}

	if err := ValidateP2PInput(p2pInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid request body",
			"errors":  err,
		})
	}

	fromUser, err := h.service.GetUser(p2pInput.FromUserID)
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

	toUser, err := h.service.GetUser(p2pInput.ToUserID)
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
	c.Locals("p2pInput", p2pInput)
	return c.Next()
}

func (h *Handler) CheckBalanceOperationInput(c *fiber.Ctx) error {
	balanceOperationInput := domain.BalanceOperationInput{}

	if err := c.BodyParser(&balanceOperationInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "parsing data from request body failed with error: " + err.Error(),
		})
	}

	if err := ValidateBalanceOperationInput(balanceOperationInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid request body",
			"errors":  err,
		})
	}

	if balanceOperationInput.Type != "add" && balanceOperationInput.Type != "subtract" {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "invalid balance operation type",
		})
	}

	user, err := h.service.GetUser(balanceOperationInput.UserID)
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
		user = &domain.User{
			ID: balanceOperationInput.UserID,
		}
		err := h.service.CreateUser(user)
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
