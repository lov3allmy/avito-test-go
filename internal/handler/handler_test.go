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
	"strconv"
	"strings"
	"testing"
)

func TestHandler_makeP2PTransfer(t *testing.T) {

	type mockBehavior func(s *mock_domain.MockService, p2pInput domain.P2PInput)

	tests := []struct {
		name                 string
		inputBody            string
		p2pInput             domain.P2PInput
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Case OK",
			inputBody: `{"from_user_id":1,"to_user_id":2,"amount":10}`,
			p2pInput: domain.P2PInput{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     10,
			},
			mockBehavior: func(s *mock_domain.MockService, p2pInput domain.P2PInput) {
				s.EXPECT().MakeP2PTransfer(p2pInput).Return(nil)
			},
			expectedStatusCode:   fiber.StatusOK,
			expectedResponseBody: `{"message":"transfer completed"}`,
		},
		{
			name:      "Case ServerInternalError",
			inputBody: `{"from_user_id":1,"to_user_id":2,"amount":10}`,
			p2pInput: domain.P2PInput{
				FromUserID: 1,
				ToUserID:   2,
				Amount:     10,
			},
			mockBehavior: func(s *mock_domain.MockService, p2pInput domain.P2PInput) {
				s.EXPECT().MakeP2PTransfer(p2pInput).Return(errors.New("service returning error"))
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
			test.mockBehavior(service, test.p2pInput)

			handler := NewHandler(service)

			app := fiber.New()
			api := app.Group("/api")
			api.Post("/p2p", func(c *fiber.Ctx) error {
				c.Locals("p2pInput", test.p2pInput)
				return c.Next()
			}, handler.MakeP2PTransfer)

			request := httptest.NewRequest("POST", "/api/p2p", strings.NewReader(test.inputBody))
			request.Header.Add("Content-Length", strconv.FormatInt(request.ContentLength, 10))
			request.Header.Add("Content-Type", "application/json")

			response, err := app.Test(request, -1)
			assert.Equal(t, err, nil)
			body, err := ioutil.ReadAll(response.Body)
			assert.Equal(t, err, nil)
			assert.Equal(t, string(body), test.expectedResponseBody)
			assert.Equal(t, response.StatusCode, test.expectedStatusCode)
		})
	}
}
