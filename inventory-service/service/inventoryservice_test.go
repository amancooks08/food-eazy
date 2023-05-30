package service

import (
	"net/http"
	"testing"

	"inventory-service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InventoryServiceTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *InventoryServiceTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	models.InitInventoryModels(db)
}

func (suite *InventoryServiceTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&models.Item{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryServiceTestSuite))
}

func (suite *InventoryServiceTestSuite) TestService_AddItem() {
	t := suite.T()
	type args struct {
		name        string
		description string
		price       float32
		quantity    uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantStatus uint32
	}{
		{
			name: "Add item successfully",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       100,
				quantity:    100,
			},
			wantErr: false,
			wantStatus: http.StatusCreated,
		},
		{
			name: "Add item with empty name",
			args: args{
				name:        "",
				description: "testitem1.desc",
				price:       100,
				quantity:    100,
			},
			wantErr: true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Add item with empty description",
			args: args{
				name:        "testitem1",
				description: "",
				price:       100,
				quantity:    100,
			},
			wantErr: true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Add item with zero price",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       0,
				quantity:    100,
			},
			wantErr: true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Add item with negative price",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       -100,
				quantity:    100,
			},
			wantErr: true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Add item with zero quantity",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       100,
				quantity:    0,
			},
			wantErr: true,
			wantStatus: http.StatusBadRequest,
		},
		{
			name: "Add item that already exists",
			args: args{
				name:        "testitem1",
				description: "testitem1.desc",
				price:       100,
				quantity:    100,
			},
			wantErr: true,
			wantStatus: http.StatusUnprocessableEntity,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status, item, err := AddItem(tt.args.name, tt.args.description, tt.args.price, tt.args.quantity)
			if tt.wantErr {
				assert.Equal(t, tt.wantStatus, status)
				assert.Error(t, err)
				assert.Nil(t, item)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, item)
				assert.Equal(t, tt.args.name, item.Name)
				assert.Equal(t, tt.args.name, item.Name)
				assert.Equal(t, tt.args.description, item.Description)
				assert.Equal(t, tt.args.price, item.Price)
				assert.Equal(t, tt.args.quantity, item.Quantity)
			}
		})
	}
}

func (suite *InventoryServiceTestSuite) TestService_GetItem() {
	t := suite.T()

	t.Run("Get item successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem2",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		status, item, err = GetItem(newItem.ID)
		assert.Equal(t, uint32(http.StatusOK), status)
		assert.NoError(t, err)
		assert.NotNil(t, item)
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		status, item, err := GetItem(0)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		status, item, err := GetItem(99999)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}


func (suite *InventoryServiceTestSuite) TestService_GetAllITems() {
	t := suite.T()

	t.Run("Get all items successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem3",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem1, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem1)

		item = &models.Item{
			Name:        "testitem4",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem2, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem2)

		status, items, err := GetAllItems()
		assert.Equal(t, uint32(http.StatusOK), status)
		assert.NoError(t, err)
		assert.NotNil(t, items)
		assert.Equal(t, 4, len(items))
	})
}

func (suite *InventoryServiceTestSuite) TestService_AddQuantity() {
	t := suite.T()

	t.Run("Add quantity successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem5",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		status, newItem, err = AddQuantity(newItem.ID, 100)
		assert.Equal(t, uint32(http.StatusOK), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)
		assert.Equal(t, 200, int(newItem.Quantity))
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		status, item, err := AddQuantity(0, 100)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		status, item, err := AddQuantity(999, 100)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}

func (suite *InventoryServiceTestSuite) TestService_LowerQuantity() {
	t := suite.T()

	t.Run("Lower quantity successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem6",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		status, newItem, err = LowerQuantity(newItem.ID, 100)
		assert.Equal(t, uint32(http.StatusOK), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)
		assert.Equal(t, 0, int(newItem.Quantity))
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		status, item, err := LowerQuantity(0, 100)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		status, item, err := LowerQuantity(999, 100)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error when quantity is less than available inventory", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem7",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		status, newItem, err = LowerQuantity(newItem.ID, 200)
		assert.Equal(t, uint32(http.StatusConflict), status)
		assert.Error(t, err)
		assert.Nil(t, newItem)
	})
}


func (suite *InventoryServiceTestSuite) TestService_DeleteItem() {
	t := suite.T()

	t.Run("Delete item successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem8",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		status, newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.Equal(t, uint32(http.StatusCreated), status)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		status, err = DeleteItem(newItem.ID)
		assert.Equal(t, uint32(http.StatusOK), status)
		assert.NoError(t, err)

		status, item, err = GetItem(newItem.ID)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		status, err := DeleteItem(0)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		status, err := DeleteItem(999)
		assert.Equal(t, uint32(http.StatusNotFound), status)
		assert.Error(t, err)
	})
}