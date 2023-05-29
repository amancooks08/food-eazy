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
		testItem := Item{
			Name:        "testitem1",
			Description: "testitem1.desc",
			Price:       100,
			Quantity:    100,
		}

		err := CreateItem(testItem)
		assert.NoError(t, err)
	})
}
