package models

import (
	"testing"

	"auth-service/errors"

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
	assert.NoError(suite.T(), err)
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

func (suite *AuthModelsTestSuite) TestModels_RegisterUser() {
	t := suite.T()

	t.Run("Register user with valid details", func(t *testing.T) {
		testUser := User{
			Name:        "test",
			Email:       "test1@gmail.com",
			Password:    "test123",
			PhoneNumber: "1234567890",
		}

		err := RegisterUser(&testUser)
		assert.NoError(t, err)
	})

	t.Run("Register user with nil user", func(t *testing.T) {
		err := RegisterUser(nil)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), errors.ErrInvalidUser.Error())
	})

	t.Run("Register user with duplicate email", func(t *testing.T) {
		testUser := User{
			Name:        "test2",
			Email:       "test2@mail.com",
			Password:    "test123",
			PhoneNumber: "1234567891",
		}

		err1 := RegisterUser(&testUser)
		assert.NoError(t, err1)

		duplicateUser := User{
			Name:        "test3",
			Email:       "test2@mail.com",
			Password:    "test123",
			PhoneNumber: "1234567893",
		}

		err2 := RegisterUser(&duplicateUser)
		assert.Error(t, err2)
		assert.Equal(t, err2.Error(), errors.ErrCreateUser.Error())
	})

	t.Run("Register user with duplicate phone number", func(t *testing.T) {
		testUser := User{
			Name:        "test3",
			Email:       "test3@mailcom",
			Password:    "test123",
			PhoneNumber: "1234567893",
		}

		err1 := RegisterUser(&testUser)
		assert.NoError(t, err1)

		duplicateUser := User{
			Name:        "test4",
			Email:       "test4@mail.com",
			Password:    "test123",
			PhoneNumber: "1234567893",
		}

		err2 := RegisterUser(&duplicateUser)
		assert.Error(t, err2)
		assert.Equal(t, err2.Error(), errors.ErrCreateUser.Error())
	})
}
