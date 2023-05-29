package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type InventoryModelsTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *InventoryModelsTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(suite.T(), err)
	suite.db = db
	InitInventoryModels(db)
}

func (suite *InventoryModelsTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&Item{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestInventoryModelsTestSuite(t *testing.T) {
	suite.Run(t, new(InventoryModelsTestSuite))
}

func (suite *InventoryModelsTestSuite) TestItemModelInitialized() {
	t := suite.T()
	t.Run("User model should be initialized", func(t *testing.T) {
		err := suite.db.AutoMigrate(&Item{})
		assert.NoError(t, err)
	})
}

func (suite *InventoryModelsTestSuite) TestModels_AddItem() {
	t := suite.T()

	t.Run("Add item successfully", func(t *testing.T) {
		testItem := &Item{
			Name:        "testitem1",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		item, err := CreateItem(testItem)
		assert.NoError(t, err)
		assert.NotNil(t, item)
		assert.Equal(t, testItem.Name, item.Name)
		assert.Equal(t, testItem.Description, item.Description)
		assert.Equal(t, testItem.Price, item.Price)
		assert.Equal(t, testItem.Quantity, item.Quantity)
	})

	t.Run("Add item with nil value", func(t *testing.T) {
		item, err := CreateItem(nil)
		assert.Error(t, err)
		assert.Nil(t, item)
	})

	t.Run("Add item with duplicate name", func(t *testing.T) {
		testItem1 := &Item{
			Name:        "testitem2",
			Description: "testitem2.desc",
			Price:       100,
			Quantity:    100,
		}

		item, err := CreateItem(testItem1)
		assert.NoError(t, err)
		assert.NotNil(t, item)

		testItem2 := &Item{
			Name:        "testitem2",
			Description: "testitem2.desc",
			Price:       100,
			Quantity:    100,
		}

		item, err = CreateItem(testItem2)
		assert.Error(t, err)
		assert.Nil(t, item)
	})
}

func (suite *InventoryModelsTestSuite) TestModels_GetItem() {
	t := suite.T()

	t.Run("Get item successfully", func(t *testing.T) {
		testItem := &Item{
			Name:        "testitem3",
			Description: "testitem3.desc",
			Price:       100,
			Quantity:    10,
		}

		item, err := CreateItem(testItem)
		assert.NoError(t, err)
		assert.NotNil(t, item)

		got, err := GetItem(item.ID)
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, testItem.Name, item.Name)
		assert.Equal(t, testItem.Description, item.Description)
		assert.Equal(t, testItem.Price, item.Price)
		assert.Equal(t, testItem.Quantity, item.Quantity)
	})

	t.Run("Get item with invalid id", func(t *testing.T) {
		got, err := GetItem(0)
		assert.Error(t, err)
		assert.Nil(t, got)
	})
}


func (suite *InventoryModelsTestSuite) TestModels_GetAllItems() {
	t := suite.T()

	t.Run("Get all items successfully", func(t *testing.T) {
		testItem1 := &Item{
			Name:        "testitem4",
			Description: "testitem4.desc",
			Price:       100,
			Quantity:    10,
		}

		item1, err := CreateItem(testItem1)
		assert.NoError(t, err)
		assert.NotNil(t, item1)

		testItem2 := &Item{
			Name:        "testitem5",
			Description: "testitem5.desc",
			Price:       100,
			Quantity:    10,
		}

		item2, err := CreateItem(testItem2)
		assert.NoError(t, err)
		assert.NotNil(t, item2)

		got, err := GetAllItems()
		assert.NoError(t, err)
		assert.NotNil(t, got)
		assert.Equal(t, 4, len(got))
	})
}