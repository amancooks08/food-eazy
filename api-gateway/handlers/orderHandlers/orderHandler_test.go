package orderHandlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api-gateway/dependencies"
	"api-gateway/domain"
	mocks "api-gateway/mocks/ordermocks"
	proto "api-gateway/proto/order"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrderHandlerTestSuite struct {
	suite.Suite
	grpc *mocks.OrderServiceClient
}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(OrderHandlerTestSuite))
}
func (suite *OrderHandlerTestSuite) SetupTest() {
	suite.grpc = &mocks.OrderServiceClient{}
}

func (suite *OrderHandlerTestSuite) TearDownTest() {
	suite.grpc.AssertExpectations(suite.T())
}

func (suite *OrderHandlerTestSuite) TestOrderHandler_PlaceOrder() {
	t := suite.T()

	t.Run("expect to return 201 when order placed successfully", func(t *testing.T) {
		// Arrange
		requestBody := domain.PlaceOrderRequest{
			ItemID:   1,
			Quantity: 2,
		}

		expectedRequest := proto.PlaceOrderRequest{
			UserId:   1,
			ItemId:   requestBody.ItemID,
			Quantity: requestBody.Quantity,
		}

		expectedResponse := proto.PlaceOrderResponse{
			StatusCode: http.StatusCreated,
			Order: &proto.Order{
				OrderId:   1,
				UserId:    1,
				ItemId:    1,
				Quantity:  2,
				Amount:    200,
				OrderTime: "2021-01-01 00:00:00",
			},
		}

		response := domain.PlaceOrderResponse{
			OrderID:   expectedResponse.Order.OrderId,
			Amount:    expectedResponse.Order.Amount,
			OrderTime: expectedResponse.Order.OrderTime,
		}

		respBody, err := json.Marshal(response)
		assert.NoError(t, err)

		reqBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest("POST", "/order/place", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()
		ctx := req.Context()
		req = req.WithContext(context.WithValue(ctx, "id", 1))

		// Act
		suite.grpc.On("PlaceOrder", req.Context(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := dependencies.Dependencies{
			OrderService: suite.grpc,
		}
		handler := PlaceOrder(deps.OrderService)
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, string(respBody), res.Body.String())
	})

	t.Run("expect to return 400 when request body is invalid", func(t *testing.T) {
		// Arrange
		requestBody := domain.PlaceOrderRequest{
			ItemID:   1,
			Quantity: 2,
		} 

		response := domain.Message{
			Message: "unauthorized access: invalid user id",
		}

		respBody, err := json.Marshal(response)
		assert.NoError(t, err)

		reqBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest("POST", "/user/order/place", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()
		ctx := req.Context()
		req = req.WithContext(context.WithValue(ctx, "id", 0))

		// Act
		deps := dependencies.Dependencies{
			OrderService: suite.grpc,
		}
		handler := PlaceOrder(deps.OrderService)
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, string(respBody), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 422 when item is not available in enough quantity", func(t *testing.T) {
		// Arrange
		requestBody := domain.PlaceOrderRequest{
			ItemID:   1,
			Quantity: 2,
		}

		expectedRequest := proto.PlaceOrderRequest{
			UserId:   1,
			ItemId:   requestBody.ItemID,
			Quantity: requestBody.Quantity,
		}

		expectedResponse := proto.PlaceOrderResponse{
			StatusCode: 422,
			Order:      nil,
		}

		response := domain.Message{
			Message: "grpc received error: item not available in enough quantity",
		}

		respBody, err := json.Marshal(response)
		assert.NoError(t, err)

		reqBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest("POST", "/order/place", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()
		ctx := req.Context()
		req = req.WithContext(context.WithValue(ctx, "id", 1))

		// Act
		suite.grpc.On("PlaceOrder", req.Context(), &expectedRequest).Return(&expectedResponse, errors.New("item not available in enough quantity")).Once()
		deps := dependencies.Dependencies{
			OrderService: suite.grpc,
		}
		handler := PlaceOrder(deps.OrderService)
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, string(respBody), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 401 when user id is not present in context", func(t *testing.T) {
		// Arrange
		requestBody := domain.PlaceOrderRequest{
			ItemID:   1,
			Quantity: 2,
		}

		response := domain.Message{
			Message: "unauthorized access: invalid user id",
		}

		respBody, err := json.Marshal(response)
		assert.NoError(t, err)

		reqBody, err := json.Marshal(requestBody)
		assert.NoError(t, err)

		jsonRequest := string(reqBody)
		req := httptest.NewRequest("POST", "/order/place", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		deps := dependencies.Dependencies{
			OrderService: suite.grpc,
		}
		handler := PlaceOrder(deps.OrderService)
		handler.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, http.StatusUnauthorized, res.Code)
		assert.Equal(t, string(respBody), strings.Split(res.Body.String(), "\n")[0])
	})
}
