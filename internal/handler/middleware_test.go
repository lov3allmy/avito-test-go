package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/lov3allmy/avito-test-go/internal/domain"
	mock_domain "github.com/lov3allmy/avito-test-go/internal/mocks"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_CheckGetBalanceInput(t *testing.T) {
	type mockBehavior func(s *mock_domain.MockService, userID int, user *domain.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputObject          domain.GetBalanceInput
		user                 domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:        "OK",
			inputBody:   `{"user_id":1}`,
			inputObject: domain.GetBalanceInput{ID: 1},
			user: domain.User{
				ID:      1,
				Balance: 0,
			},
			mockBehavior: func(s *mock_domain.MockService, userID int, user *domain.User) {
				s.EXPECT().GetUser(userID).Return(user, nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"ok"}`,
		},
		{
			name:                 "Invalid user_id",
			inputBody:            `{"user_id":-1}`,
			inputObject:          domain.GetBalanceInput{},
			user:                 domain.User{},
			mockBehavior:         func(s *mock_domain.MockService, userID int, user *domain.User) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"errors":[{"FailedField":"GetBalanceInput.ID","Tag":"min","Value":"0"}],"message":"invalid request body"}`,
		},
		{
			name:                 "Required user_id",
			inputBody:            `{}`,
			inputObject:          domain.GetBalanceInput{},
			user:                 domain.User{},
			mockBehavior:         func(s *mock_domain.MockService, userID int, user *domain.User) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"errors":[{"FailedField":"GetBalanceInput.ID","Tag":"required","Value":""}],"message":"invalid request body"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_domain.NewMockService(c)
			test.mockBehavior(service, test.inputObject.ID, &test.user)

			handler := NewHandler(service)

			app := fiber.New()
			app.Get("", handler.CheckGetBalanceInput, func(ctx *fiber.Ctx) error {
				assert.Equal(t, ctx.Locals("user").(*domain.User), &test.user)
				return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
					"message": "ok",
				})
			})

			request := httptest.NewRequest("GET", "/", strings.NewReader(test.inputBody))
			request.Header.Add("Content-Type", "application/json")

			response, err := app.Test(request)
			assert.Equal(t, err, nil)

			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)

			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}

func TestHandler_CheckP2PInput(t *testing.T) {
	type mockBehavior func(s *mock_domain.MockService, input domain.P2PInput, fromUser, toUser *domain.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputObject          domain.P2PInput
		fromUser             domain.User
		toUser               domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"from_user_id":1,"to_user_id":2,"amount":10}`,
			inputObject: domain.P2PInput{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     10,
			},
			fromUser: domain.User{
				ID:      1,
				Balance: 10,
			},
			toUser: domain.User{
				ID:      2,
				Balance: 0,
			},
			mockBehavior: func(s *mock_domain.MockService, input domain.P2PInput, fromUser, toUser *domain.User) {
				s.EXPECT().GetUser(input.FromUserID).Return(fromUser, nil)
				s.EXPECT().GetUser(input.ToUserID).Return(toUser, nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"ok"}`,
		},
		{
			name:      "Too much amount",
			inputBody: `{"from_user_id":1,"to_user_id":2,"amount":100}`,
			inputObject: domain.P2PInput{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     100,
			},
			fromUser: domain.User{
				ID:      1,
				Balance: 10,
			},
			toUser: domain.User{
				ID:      2,
				Balance: 0,
			},
			mockBehavior: func(s *mock_domain.MockService, input domain.P2PInput, fromUser, toUser *domain.User) {
				s.EXPECT().GetUser(input.FromUserID).Return(fromUser, nil)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"message":"not enough balance to make transfer"}`,
		},
		{
			name:                 "Invalid input 1",
			inputBody:            `{"from_user_id":1,"to_user_id":1,"amount":10}`,
			inputObject:          domain.P2PInput{},
			fromUser:             domain.User{},
			toUser:               domain.User{},
			mockBehavior:         func(s *mock_domain.MockService, input domain.P2PInput, fromUser, toUser *domain.User) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"errors":[{"FailedField":"ToUserID","Tag":"nefield","Value":"FromUserID"}],"message":"invalid request body"}`,
		},
		{
			name:                 "Invalid input 2",
			inputBody:            `{"from_user_id":-1,"to_user_id":1,"amount":10}`,
			inputObject:          domain.P2PInput{},
			fromUser:             domain.User{},
			toUser:               domain.User{},
			mockBehavior:         func(s *mock_domain.MockService, input domain.P2PInput, fromUser, toUser *domain.User) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"errors":[{"FailedField":"FromUserID","Tag":"min","Value":"0"}],"message":"invalid request body"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_domain.NewMockService(c)
			test.mockBehavior(service, test.inputObject, &test.fromUser, &test.toUser)

			handler := NewHandler(service)

			app := fiber.New()
			app.Post("", handler.CheckP2PInput, func(ctx *fiber.Ctx) error {
				assert.Equal(t, ctx.Locals("p2pInput").(domain.P2PInput), test.inputObject)
				return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
					"message": "ok",
				})
			})

			request := httptest.NewRequest("POST", "/", strings.NewReader(test.inputBody))
			request.Header.Add("Content-Type", "application/json")

			response, err := app.Test(request)
			assert.Equal(t, err, nil)

			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)

			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}

func TestHandler_CheckBalanceOperationInput(t *testing.T) {
	type mockBehavior func(s *mock_domain.MockService, userID int, user *domain.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputObject          domain.BalanceOperationInput
		user                 domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"user_id":1,"amount":10,"type":"add"}`,
			inputObject: domain.BalanceOperationInput{
				UserID: 1,
				Amount: 10,
				Type:   "add",
			},
			user: domain.User{
				ID:      1,
				Balance: 0,
			},
			mockBehavior: func(s *mock_domain.MockService, userID int, user *domain.User) {
				s.EXPECT().GetUser(userID).Return(user, nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"ok"}`,
		},
		{
			name:                 "Invalid type in input",
			inputBody:            `{"user_id":1,"amount":10,"type":"addd"}`,
			inputObject:          domain.BalanceOperationInput{},
			user:                 domain.User{},
			mockBehavior:         func(s *mock_domain.MockService, userID int, user *domain.User) {},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"errors":[{"FailedField":"BalanceOperationInput.Type","Tag":"oneof","Value":"add subtract"}],"message":"invalid request body"}`,
		},
		{
			name:      "User not found",
			inputBody: `{"user_id":1,"amount":10,"type":"subtract"}`,
			inputObject: domain.BalanceOperationInput{
				UserID: 1,
				Amount: 10,
				Type:   "subtract",
			},
			user: domain.User{},
			mockBehavior: func(s *mock_domain.MockService, userID int, user *domain.User) {
				s.EXPECT().GetUser(userID).Return(nil, nil)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"message":"there is no user with that \"user_id\""}`,
		},
		{
			name:      "Too much amount",
			inputBody: `{"user_id":1,"amount":10,"type":"subtract"}`,
			inputObject: domain.BalanceOperationInput{
				UserID: 1,
				Amount: 10,
				Type:   "subtract",
			},
			user: domain.User{
				ID:      1,
				Balance: 0,
			},
			mockBehavior: func(s *mock_domain.MockService, userID int, user *domain.User) {
				s.EXPECT().GetUser(userID).Return(user, nil)
			},
			expectedStatusCode:   fiber.StatusBadRequest,
			expectedResponseBody: `{"message":"not enough balance to make operation"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_domain.NewMockService(c)
			test.mockBehavior(service, test.inputObject.UserID, &test.user)

			handler := NewHandler(service)

			app := fiber.New()
			app.Post("", handler.CheckBalanceOperationInput, func(ctx *fiber.Ctx) error {
				assert.Equal(t, ctx.Locals("user").(*domain.User), &test.user)
				assert.Equal(t, ctx.Locals("operationAmount").(int), test.inputObject.Amount)
				assert.Equal(t, ctx.Locals("operationType").(string), test.inputObject.Type)
				return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
					"message": "ok",
				})
			})

			request := httptest.NewRequest("POST", "/", strings.NewReader(test.inputBody))
			request.Header.Add("Content-Type", "application/json")

			response, err := app.Test(request)
			assert.Equal(t, err, nil)

			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)

			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}
