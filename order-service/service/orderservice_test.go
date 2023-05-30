package service

import (
	"context"
	"errors"
	"net/http"
	mocks "order-service/mocks/inventoryMocks"
	"order-service/models"
	proto "order-service/proto/inventorypb"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderServiceTestSuite struct {
	suite.Suite
	db      *gorm.DB
	service *OrderService
}

func (suite *OrderServiceTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	models.InitOrderModels(db)

	suite.service = &OrderService{
		client: &mocks.InventoryServiceClient{},
	}
}

func (suite *OrderServiceTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&models.Order{})
	sql, _ := suite.db.DB()
	sql.Close()

	suite.service = nil
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (suite *OrderServiceTestSuite) TestService_PlaceOrder() {
	t := suite.T()

	t.Run("PlaceOrder successfully", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(1)
		quantity, amount := uint32(2), float32(20.0)

		itemResponse := &proto.GetItemResponse{
			StatusCode:  http.StatusOK,
			Id:          itemID,
			Name:        "item",
			Description: "item description",
			Price:       10,
			Quantity:    12,
		}

		order := &models.Order{
			UserID:    userID,
			ItemID:    itemID,
			Quantity:  quantity,
			Amount:    amount,
			OrderTime: "2021-01-01 00:00:00",
		}

		//Act
		suite.service.client.(*mocks.InventoryServiceClient).On("GetItem", context.Background(), &proto.GetItemRequest{
			Id: itemID}).Return(itemResponse, nil).Once()

		suite.service.client.(*mocks.InventoryServiceClient).On("LowerQuantity", context.Background(), &proto.LowerQuantityRequest{
			Id:       itemID,
			Quantity: quantity,
		}).Return(&proto.LowerQuantityResponse{
			StatusCode: http.StatusOK,
			Id:         itemID,
			Quantity:   itemResponse.Quantity - quantity,
		}, nil).Once()

		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.Equal(t, order.ItemID, placedOrder.ItemID)
		assert.Equal(t, order.Quantity, placedOrder.Quantity)
		assert.Equal(t, order.Amount, placedOrder.Amount)
		assert.Equal(t, order.OrderTime, "2021-01-01 00:00:00")
	})

	t.Run("PlaceOrder failed due to item not found", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(1)
		quantity := uint32(2)

		itemResponse := &proto.GetItemResponse{
			StatusCode:  http.StatusNotFound,
			Id:          0,
			Name:        "",
			Description: "",
			Price:       0,
			Quantity:    0,
		}

		//Act
		suite.service.client.(*mocks.InventoryServiceClient).On("GetItem", context.Background(), &proto.GetItemRequest{
			Id: itemID}).Return(itemResponse, errors.New("mocked error")).Once()

		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Nil(t, placedOrder)
	})

	t.Run("PlaceOrder failed due to item quantity not enough", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(1)
		quantity := uint32(2)

		itemResponse := &proto.GetItemResponse{
			StatusCode:  http.StatusOK,
			Id:          itemID,
			Name:        "item",
			Description: "item description",
			Price:       10,
			Quantity:    1,
		}

		//Act
		suite.service.client.(*mocks.InventoryServiceClient).On("GetItem", context.Background(), &proto.GetItemRequest{
			Id: itemID}).Return(itemResponse, nil).Once()

		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, uint32(http.StatusUnprocessableEntity), status)
		assert.Nil(t, placedOrder)
	})

	t.Run("PlaceOrder failed due to zero quantity", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(1)
		quantity := uint32(0)

		//Act
		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, uint32(http.StatusBadRequest), status)
		assert.Nil(t, placedOrder)
	})

	t.Run("PlaceOrder failed due to zero userID", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(0)
		quantity := uint32(1)

		//Act
		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, uint32(http.StatusBadRequest), status)
		assert.Nil(t, placedOrder)
	})

	t.Run("PlaceOrder failed due to zero itemID", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(0), uint32(1)
		quantity := uint32(1)

		//Act
		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, uint32(http.StatusBadRequest), status)
		assert.Nil(t, placedOrder)
	})

	t.Run("PlaceOrder failed when LowerQuantity returned error", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(1)
		quantity := uint32(2)

		itemResponse := &proto.GetItemResponse{
			StatusCode:  http.StatusOK,
			Id:          itemID,
			Name:        "item",
			Description: "item description",
			Price:       10,
			Quantity:    12,
		}

		//Act
		suite.service.client.(*mocks.InventoryServiceClient).On("GetItem", context.Background(), &proto.GetItemRequest{
			Id: itemID}).Return(itemResponse, nil).Once()

		suite.service.client.(*mocks.InventoryServiceClient).On("LowerQuantity", context.Background(), &proto.LowerQuantityRequest{
			Id:       itemID,
			Quantity: quantity,
		}).Return(&proto.LowerQuantityResponse{
			StatusCode: http.StatusInternalServerError,
			Id:         0,
			Quantity:   0,
		}, errors.New("mocked error")).Once()

		status, placedOrder, err := suite.service.PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, uint32(http.StatusInternalServerError), status)
		assert.Nil(t, placedOrder)
	})
}
