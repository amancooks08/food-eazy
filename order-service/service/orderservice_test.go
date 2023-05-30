package service

import (
	"context"
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
	db         *gorm.DB
	mockclient *mocks.InventoryServiceClient
}

func (suite *OrderServiceTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	models.InitOrderModels(db)

	suite.mockclient = &mocks.InventoryServiceClient{}
}

func (suite *OrderServiceTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&models.Order{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (suite *OrderServiceTestSuite) TestService_PlaceOrder() {
	t := suite.T()

	t.Run("PlaceOrder successfully", func(t *testing.T) {
		//Arrange
		itemID, userID := uint32(1), uint32(1)
		quantity, amount := uint32(2), float32(1.0)

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
		suite.mockclient.On("GetItem", context.Background(), &proto.GetItemRequest{
			Id: itemID,
		}).Return(itemResponse, nil).Once()

		suite.mockclient.On("LowerQuantity", context.Background(), &proto.LowerQuantityRequest{
			Id:       itemID,
			Quantity: quantity}).Return(&proto.LowerQuantityResponse{
			StatusCode: http.StatusOK,
			Id:         1,
			Quantity:   10}, nil).Once()

		status, placedOrder, err := PlaceOrder(userID, itemID, quantity)

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, status)
		assert.Equal(t, order.ID, placedOrder.ID)
		assert.Equal(t, order.UserID, placedOrder.UserID)
		assert.Equal(t, order.ItemID, placedOrder.ItemID)
		assert.Equal(t, order.Quantity, placedOrder.Quantity)
		assert.Equal(t, order.Amount, placedOrder.Amount)
		assert.Equal(t, order.OrderTime, placedOrder.OrderTime)
	})
}
