package service

import (
	"testing"

	"inventory-service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthServiceTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *AuthServiceTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	models.InitInventoryModels(db)
}

func (suite *AuthServiceTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&models.Item{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (suite *AuthServiceTestSuite) TestService_AddItem() {
	t := suite.T()
	type args struct {
		name        string
		description string
		price       float64
		quantity    uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			item, err := AddItem(tt.args.name, tt.args.description, tt.args.price, tt.args.quantity)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, item)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, item)
				assert.Equal(t, tt.args.name, item.Name)
				assert.Equal(t, tt.args.description, item.Description)
				assert.Equal(t, tt.args.price, item.Price)
				assert.Equal(t, tt.args.quantity, item.Quantity)
			}
		})
	}
}

func (suite *AuthServiceTestSuite) TestService_GetItem() {
	t := suite.T()

	t.Run("Get item successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem2",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		item, err = GetItem(newItem.ID)
		assert.NoError(t, err)
		assert.NotNil(t, item)
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		item, err := GetItem(0)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		item, err := GetItem(999)
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}


func (suite *AuthServiceTestSuite) TestService_GetAllITems() {
	t := suite.T()

	t.Run("Get all items successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem3",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		newItem1, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.NoError(t, err)
		assert.NotNil(t, newItem1)

		item = &models.Item{
			Name:        "testitem4",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		newItem2, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.NoError(t, err)
		assert.NotNil(t, newItem2)

		items, err := GetAllItems()
		assert.NoError(t, err)
		assert.NotNil(t, items)
		assert.Equal(t, 4, len(items))
	})
}

func (suite *AuthServiceTestSuite) TestService_AddQuantity() {
	t := suite.T()

	t.Run("Add quantity successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem5",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		newItem, err = AddQuantity(newItem.ID, 100)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)
		assert.Equal(t, 200, int(newItem.Quantity))
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		item, err := AddQuantity(0, 100)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		item, err := AddQuantity(999, 100)
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}

func (suite *AuthServiceTestSuite) TestService_LowerQuantity() {
	t := suite.T()

	t.Run("Lower quantity successfully", func(t *testing.T) {
		item := &models.Item{
			Name:        "testitem6",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		newItem, err = LowerQuantity(newItem.ID, 100)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)
		assert.Equal(t, 0, int(newItem.Quantity))
	})

	t.Run("expect error with invalid id", func(t *testing.T) {
		item, err := LowerQuantity(0, 100)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("expect error with non-exist id", func(t *testing.T) {
		item, err := LowerQuantity(999, 100)
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

		newItem, err := AddItem(item.Name, item.Description, item.Price, item.Quantity)
		assert.NoError(t, err)
		assert.NotNil(t, newItem)

		newItem, err = LowerQuantity(newItem.ID, 200)
		assert.Error(t, err)
		assert.Nil(t, newItem)
	})
}

