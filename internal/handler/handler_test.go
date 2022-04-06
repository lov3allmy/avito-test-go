package handler

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/lov3allmy/avito-test-go/internal/domain"
	mock_domain "github.com/lov3allmy/avito-test-go/internal/mocks"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandler_makeP2PTransfer(t *testing.T) {

	type mockBehavior func(s *mock_domain.MockService, input domain.P2PInput)

	tests := []struct {
		name                 string
		inputObject          domain.P2PInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputObject: domain.P2PInput{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     10,
			},
			mockBehavior: func(s *mock_domain.MockService, input domain.P2PInput) {
				s.EXPECT().MakeP2PTransfer(input).Return(nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"transfer completed"}`,
		},
		{
			name: "InternalServerError",
			inputObject: domain.P2PInput{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     10,
			},
			mockBehavior: func(s *mock_domain.MockService, input domain.P2PInput) {
				s.EXPECT().MakeP2PTransfer(input).Return(errors.New("service returning error"))
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: `{"message":"making transfer failed with error: service returning error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_domain.NewMockService(c)
			test.mockBehavior(service, test.inputObject)

			handler := NewHandler(service)

			app := fiber.New()
			app.Post("", func(ctx *fiber.Ctx) error {
				ctx.Locals("p2pInput", test.inputObject)
				return ctx.Next()
			}, handler.MakeP2PTransfer)

			request := httptest.NewRequest("POST", "/", nil)

			response, err := app.Test(request)
			assert.Equal(t, err, nil)

			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)

			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}

func TestHandler_GetBalanceByUserID(t *testing.T) {

	tests := []struct {
		name                 string
		inputObject          domain.User
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "OK",
			inputObject:          domain.User{ID: 1, Balance: 10},
			expectedStatusCode:   200,
			expectedResponseBody: `{"balance":10}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_domain.NewMockService(c)

			handler := NewHandler(service)

			app := fiber.New()
			app.Get("", func(ctx *fiber.Ctx) error {
				ctx.Locals("user", &test.inputObject)
				return ctx.Next()
			}, handler.GetBalanceByUserID)

			request := httptest.NewRequest("GET", "/", nil)

			response, err := app.Test(request)
			assert.Equal(t, err, nil)

			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)

			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}

func TestHandler_MakeBalanceOperationByUserID(t *testing.T) {

	type mockBehavior func(s *mock_domain.MockService, userID int, user *domain.User)

	tests := []struct {
		name                 string
		inputObject          domain.BalanceOperationInput
		user                 domain.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			inputObject: domain.BalanceOperationInput{
				UserID: 1,
				Amount: 10,
				Type:   "add",
			},
			user: domain.User{
				ID:      1,
				Balance: 10,
			},
			mockBehavior: func(s *mock_domain.MockService, userID int, user *domain.User) {
				s.EXPECT().UpdateUser(userID, user).Return(nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"operation completed"}`,
		},
		{
			name: "InternalServerError",
			inputObject: domain.BalanceOperationInput{
				UserID: 1,
				Amount: 10,
				Type:   "add",
			},
			mockBehavior: func(s *mock_domain.MockService, userID int, user *domain.User) {
				s.EXPECT().UpdateUser(userID, user).Return(errors.New("service returning error"))
			},
			expectedStatusCode:   fiber.StatusInternalServerError,
			expectedResponseBody: `{"message":"making operation failed with error: service returning error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_domain.NewMockService(c)
			test.mockBehavior(service, test.user.ID, &test.user)

			handler := NewHandler(service)

			app := fiber.New()
			app.Post("", func(ctx *fiber.Ctx) error {
				ctx.Locals("user", &test.user)
				ctx.Locals("operationAmount", test.inputObject.Amount)
				ctx.Locals("operationType", test.inputObject.Type)
				return ctx.Next()
			}, handler.MakeBalanceOperationByUserID)

			request := httptest.NewRequest("POST", "/", nil)

			response, err := app.Test(request)
			assert.Equal(t, err, nil)

			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)

			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}
