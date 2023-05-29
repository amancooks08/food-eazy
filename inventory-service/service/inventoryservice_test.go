package service

import (
	"testing"

	"inventory-service/models"
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
