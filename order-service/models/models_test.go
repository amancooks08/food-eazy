package models

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderModelTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *OrderModelTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	InitOrderModels(db)
}

func (suite *OrderModelTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&Order{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestOrderModelTestSuite(t *testing.T) {
	suite.Run(t, new(OrderModelTestSuite))
}

func (suite *OrderModelTestSuite) TestOrderModel_CreateOrder() {
	t := suite.T()

	type args struct {
		order *Order
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantStatus uint32
	}{
		{
			name: "CreateOrder successfully",
			args: args{
				order: &Order{
					UserID:   1,
					ItemID:   1,
					Quantity: 1,
					Amount:   1.0,
				},
			},
			wantErr:    false,
			wantStatus: http.StatusCreated,
		},
		{
			name: "CreateOrder failed due to nil order",
			args: args{
				order: nil,
			},
			wantErr:    true,
			wantStatus: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStatus, err := CreateOrder(tt.args.order)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.wantStatus, gotStatus)
		})
	}
}

func (suite *OrderModelTestSuite) TestOrderModel_GetOrder() {
	t := suite.T()

	t.Run("GetOrder successfully", func(t *testing.T) {
		//Arrange
		orderID := uint32(1)
		order := &Order{
			UserID:   1,
			ItemID:   1,
			Quantity: 1,
			Amount:   1.0,
		}
		suite.db.Create(order)

		//Act
		gotStatus, gotOrder, err := GetOrder(orderID)

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, int(gotStatus))
		assert.Equal(t, order.Amount, gotOrder.Amount)
		assert.Equal(t, order.ItemID, gotOrder.ItemID)
		assert.Equal(t, order.Quantity, gotOrder.Quantity)
		assert.Equal(t, order.UserID, gotOrder.UserID)
		assert.Equal(t, orderID, gotOrder.ID)

	})

	t.Run("GetOrder failed due to invalid orderID", func(t *testing.T) {
		//Arrange
		orderID := uint32(0)

		//Act
		gotStatus, gotOrder, err := GetOrder(orderID)

		//Assert
		assert.Error(t, err)
		assert.Equal(t, http.StatusBadRequest, int(gotStatus))
		assert.Nil(t, gotOrder)
	})
} 