package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type AuthModelsTestSuite struct {
	suite.Suite
	db *gorm.DB
}

func (suite *AuthModelsTestSuite) SetupSuite() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.FailNow("failed to connect database")
	}
	suite.db = db
	InitAuthModels(db)
}

func (suite *AuthModelsTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&User{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestAuthModelsTestSuite(t *testing.T) {
	suite.Run(t, new(AuthModelsTestSuite))
}

func (suite *AuthModelsTestSuite) TestUserModelInitialized() {
	t := suite.T()
	t.Run("User model should be initialized", func(t *testing.T) {
		err := suite.db.AutoMigrate(&User{})
		assert.NoError(t, err)
	})
}

func (suite *AuthModelsTestSuite) TestUserModel() {
	
}