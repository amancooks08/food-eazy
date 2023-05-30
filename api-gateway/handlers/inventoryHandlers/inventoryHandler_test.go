package inventoryHandlers

import (
	"api-gateway/dependencies"
	"api-gateway/domain"
	mocks "api-gateway/mocks/inventorymocks"
	proto "api-gateway/proto/inventory"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InventoryHandlersTestSuite struct {
	suite.Suite
	grpc *mocks.InventoryServiceClient
}

func (suite *InventoryHandlersTestSuite) SetupTest() {
	suite.grpc = &mocks.InventoryServiceClient{}
}

func (suite *InventoryHandlersTestSuite) TearDownTest() {
	suite.grpc.AssertExpectations(suite.T())
}

func TestInventoryHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryHandlersTestSuite))
}

func (suite *InventoryHandlersTestSuite) TestInventoryHandler_AddItem() {
	t := suite.T()

	t.Run("expect to return 201 when item added successfully", func(t *testing.T) {
		// Arrange

		requestBody := domain.AddItemRequest{
			Name:        "test1",
			Description: "test1",
			Price:       100,
			Quantity:    10,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusCreated,
			Id:          1,
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Price:       requestBody.Price,
			Quantity:    requestBody.Quantity,
		}

		response := domain.AddItemResponse{
			ID:          1,
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Price:       requestBody.Price,
			Quantity:    requestBody.Quantity,
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})

	t.Run("expect to return 400 when item name is empty", func(t *testing.T) {
		// Arrange
		requestBody := domain.AddItemRequest{
			Name:        "",
			Description: "test1",
			Price:       100,
			Quantity:    10,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusBadRequest,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: empty field",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("empty field")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 400 when item description is empty", func(t *testing.T) {
		// Arrange
		requestBody := domain.AddItemRequest{
			Name:        "test1",
			Description: "",
			Price:       100,
			Quantity:    10,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusBadRequest,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: empty field",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("empty field")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 400 when item price is 0", func(t *testing.T) {
		// Arrange
		requestBody := domain.AddItemRequest{
			Name:        "test1",
			Description: "test1",
			Price:       0,
			Quantity:    10,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusBadRequest,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: empty field",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("empty field")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 400 when item quantity is 0", func(t *testing.T) {
		// Arrange
		requestBody := domain.AddItemRequest{
			Name:        "test1",
			Description: "test1",
			Price:       100,
			Quantity:    0,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusBadRequest,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: empty field",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("empty field")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 400 when item price is negative", func(t *testing.T) {
		// Arrange
		requestBody := domain.AddItemRequest{
			Name:        "test1",
			Description: "test1",
			Price:       -100,
			Quantity:    10,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusBadRequest,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: empty field",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("empty field")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 422 when Item already exists", func(t *testing.T) {
		// Arrange
		requestBody := domain.AddItemRequest{
			Name:        "test1",
			Description: "test1",
			Price:       100,
			Quantity:    10,
		}

		expectedRequest := proto.AddItemRequest{
			Name:        requestBody.Name,
			Description: requestBody.Description,
			Quantity:    requestBody.Quantity,
			Price:       requestBody.Price,
		}

		expectedResponse := proto.AddItemResponse{
			StatusCode:  http.StatusUnprocessableEntity,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: Item already exists",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("POST", "/admin/inventory/item/add", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("AddItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("Item already exists")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusUnprocessableEntity, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 405 when method is not POST", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("GET", "/admin/inventory/item/add", nil)
		res := httptest.NewRecorder()

		// Act
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := AddItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
	})
}

func (suite *InventoryHandlersTestSuite) TestInventoryHandler_GetItem() {
	t := suite.T()

	t.Run("expect to return 200 when item is found", func(t *testing.T) {
		// Arrange
		requestBody := domain.GetItemRequest{
			ID: 1,
		}

		expectedRequest := proto.GetItemRequest{
			Id: requestBody.ID,
		}

		expectedResponse := proto.GetItemResponse{
			StatusCode:  http.StatusOK,
			Id:          1,
			Name:        "test1",
			Description: "test1",
			Price:       100,
			Quantity:    10,
		}

		response := domain.GetItemResponse{
			ID:          expectedResponse.Id,
			Name:        expectedResponse.Name,
			Description: expectedResponse.Description,
			Price:       expectedResponse.Price,
			Quantity:    expectedResponse.Quantity,
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("GET", "/admin/inventory/item/get", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("GetItem", context.Background(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 404 when item is not found", func(t *testing.T) {
		// Arrange
		requestBody := domain.GetItemRequest{
			ID: 1,
		}

		expectedRequest := proto.GetItemRequest{
			Id: requestBody.ID,
		}

		expectedResponse := proto.GetItemResponse{
			StatusCode:  http.StatusNotFound,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: Item not found",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("GET", "/admin/inventory/item/get", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("GetItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("Item not found")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusNotFound, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 400 when request body is empty", func(t *testing.T) {
		// Arrange
		requestBody := domain.GetItemRequest{}

		expectedRequest := proto.GetItemRequest{
			Id: requestBody.ID,
		}

		expectedResponse := proto.GetItemResponse{
			StatusCode:  http.StatusBadRequest,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		response := domain.Message{
			Message: "grpc received error: Invalid request",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("GET", "/admin/inventory/item/get", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("GetItem", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("Invalid request")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 405 when request method is not GET", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("POST", "/admin/inventory/item/get", nil)
		res := httptest.NewRecorder()

		// Act
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetItem(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
	})
}

func (suite *InventoryHandlersTestSuite) TestInventoryHandler_GetAllItems() {
	t := suite.T()

	t.Run("expect to return 200 when items are found", func(t *testing.T) {
		// Arrange
		expectedRequest := proto.GetAllItemsRequest{}

		expectedResponse := proto.GetAllItemsResponse{
			StatusCode: http.StatusOK,
			Items: []*proto.GetItemResponse{
				{
					StatusCode:  http.StatusOK,
					Id:          1,
					Name:        "test1",
					Description: "test1",
					Price:       100,
					Quantity:    10,
				},
				{
					StatusCode:  http.StatusOK,
					Id:          2,
					Name:        "test2",
					Description: "test2",
					Price:       200,
					Quantity:    20,
				},
			},
		}

		response := domain.GetAllItemsResponse{
			Items: []domain.GetItemResponse{
				{
					ID:          1,
					Name:        "test1",
					Description: "test1",
					Price:       100,
					Quantity:    10,
				},
				{
					ID:          2,
					Name:        "test2",
					Description: "test2",
					Price:       200,
					Quantity:    20,
				},
			},
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("GET", "/admin/inventory/items/get", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("GetAllItems", context.Background(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetAllItems(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 204 when items are not found", func(t *testing.T) {
		// Arrange
		expectedRequest := proto.GetAllItemsRequest{}

		expectedResponse := proto.GetAllItemsResponse{
			StatusCode: http.StatusNoContent,
			Items:      []*proto.GetItemResponse{},
		}

		response := domain.Message{
			Message: "grpc received error: Items not found",
		}

		exp, err := json.Marshal(response)
		assert.NoError(t, err)

		expectedReq, err := json.Marshal(expectedRequest)
		assert.NoError(t, err)
		jsonRequest := string(expectedReq)
		req := httptest.NewRequest("GET", "/admin/inventory/items/get", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act
		suite.grpc.On("GetAllItems", context.Background(), &expectedRequest).Return(&expectedResponse, errors.New("Items not found")).Once()
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetAllItems(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusNoContent, res.Code)
		assert.Equal(t, string(exp), strings.Split(res.Body.String(), "\n")[0])
	})

	t.Run("expect to return 405 when request method is not GET", func(t *testing.T) {
		// Arrange
		req := httptest.NewRequest("POST", "/admin/inventory/items/get", nil)
		res := httptest.NewRecorder()

		// Act
		deps := dependencies.Dependencies{
			InventoryService: suite.grpc,
		}
		// Assert
		handler := GetAllItems(deps.InventoryService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusMethodNotAllowed, res.Code)
	})
}
