package service

import (
	"net/http"
	"order-service/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrderServiceTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *OrderServiceTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	models.InitOrderModels(db)
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
		

}
