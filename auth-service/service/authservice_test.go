package service

import (
	"auth-service/errors"
	"auth-service/models"
	"net/http"
	"testing"

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
	models.InitAuthModels(db)
}

func (suite *AuthServiceTestSuite) TearDownSuite() {
	_ = suite.db.Migrator().DropTable(&models.User{})
	sql, _ := suite.db.DB()
	sql.Close()
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (suite *AuthServiceTestSuite) TestRegisterUser() {
	t := suite.T()
	type args struct {
		name        string
		email       string
		password    string
		phoneNumber string
		role        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Register user with valid details",
			args: args{
				name:        "test",
				email:       "test.user@gmail.com",
				password:    "test1234",
				phoneNumber: "9234567891",
				role:        "USER",
			},
			wantErr: false,
		},
		{
			name: "Register user with invalid email",
			args: args{
				name:        "test",
				email:       "test.user",
				password:    "test123",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with invalid phone number",
			args: args{
				name:        "test",
				email:       "test@mail.com",
				password:    "test123",
				phoneNumber: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Register user with short password",
			args: args{
				name:        "test",
				email:       "test@gmail.com",
				password:    "test",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Reigster user with empty name",
			args: args{
				name:        "",
				email:       "tesr@mail.com",
				password:    "test1234",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with empty email",
			args: args{
				name:        "test",
				email:       "",
				password:    "test1234",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with empty password",
			args: args{
				name:        "test",
				email:       "test@mail.com",
				password:    "",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with empty phone number",
			args: args{
				name:        "test",
				email:       "test@mail.com",
				password:    "test1234",
				phoneNumber: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := RegisterUser(tt.args.name, tt.args.email, tt.args.password, tt.args.phoneNumber, tt.args.role)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func (suite *AuthServiceTestSuite) TestLoginUser() {
	t := suite.T()

	t.Run("Login user with valid credentials", func(t *testing.T) {
		testUser := models.User{
			Name:        "testing",
			Email:       "test1@mail.com",
			Password:    "test1234",
			PhoneNumber: "1234567890",
			Role:        "USER",
		}

		err := RegisterUser(testUser.Name, testUser.Email, testUser.Password, testUser.PhoneNumber, testUser.Role)
		assert.NoError(t, err)

		token, err := LoginUser(testUser.Email, testUser.Password)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Login user with invalid password", func(t *testing.T) {
		testUser := models.User{
			Name:        "testing",
			Email:       "test2@mail.com",
			Password:    "test1234",
			PhoneNumber: "1234562890",
			Role:        "USER",
		}

		err := RegisterUser(testUser.Name, testUser.Email, testUser.Password, testUser.PhoneNumber, testUser.Role)
		assert.NoError(t, err)

		token, err := LoginUser(testUser.Email, "test123")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, errors.ErrInvalidPassword.Error(), err.Error())
	})

	t.Run("Login user with unregistered email", func(t *testing.T) {
		token, err := LoginUser("testttt@mail.com", "test1234")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, errors.ErrUserNotFound.Error(), err.Error())
	})

	t.Run("Login user with empty email", func(t *testing.T) {
		token, err := LoginUser("", "test1234")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, errors.ErrEmptyField.Error(), err.Error())
	})

	t.Run("Login user with empty password", func(t *testing.T) {
		token, err := LoginUser("test1@gmail.com", "")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, errors.ErrEmptyField.Error(), err.Error())
	})

	t.Run("Login user with invalid email", func(t *testing.T) {
		token, err := LoginUser("test1gmail.com", "test1234")
		assert.Error(t, err)
		assert.Empty(t, token)
		assert.Equal(t, errors.ErrInvalidEmail.Error(), err.Error())
	})
}

func (suite *AuthServiceTestSuite) TestValidateUser() {
	t := suite.T()

	t.Run("Validate user with valid token", func(t *testing.T) {
		testUser := models.User{
			Name:        "testing1",
			Email:       "testing1@mail.com",
			Password:    "test1234",
			PhoneNumber: "9234567898",
			Role:        "USER",
		}
		err := RegisterUser(testUser.Name, testUser.Email, testUser.Password, testUser.PhoneNumber, testUser.Role)
		assert.NoError(t, err)
		token, err := LoginUser(testUser.Email, testUser.Password)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		statusCode, err := ValidateUser(token, testUser.Role)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.NoError(t, err)
	})

	t.Run("Validate user with invalid token", func(t *testing.T) {
		statusCode, err := ValidateUser("invalid.token", "USER")
		assert.Equal(t, http.StatusUnauthorized, statusCode)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidToken.Error(), err.Error())
	})

	t.Run("Validate user with invalid role", func(t *testing.T) {
		testUser := models.User{
			Name:        "testing58",
			Email:       "testing144@mail.com",
			Password:    "test1234",
			PhoneNumber: "9234567804",
			Role:        "USER",
		}
		err := RegisterUser(testUser.Name, testUser.Email, testUser.Password, testUser.PhoneNumber, testUser.Role)
		assert.NoError(t, err)
		token, err := LoginUser(testUser.Email, testUser.Password)
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
		statusCode, err := ValidateUser(token, "TEST")
		assert.Equal(t, http.StatusUnauthorized, statusCode)
		assert.Error(t, err)
		assert.Equal(t, errors.ErrInvalidRole.Error(), err.Error())
	})
}
