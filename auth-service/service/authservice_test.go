package service

import (
	"auth-service/models"
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
		name 	  		 string
		email 	  		 string
		password  		 string
		phoneNumber 	 string
		role 	  		 string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Register user with valid details",
			args: args{
				name: "test",
				email: "test.user@gmail.com",
				password: "test1234",
				phoneNumber: "1234567890",
				role: "USER",
			},
			wantErr: false,
		},
		{
			name: "Register user with invalid email",
			args: args{
				name: "test",
				email: "test.user",
				password: "test123",
				phoneNumber: "1234567890",
			},
			wantErr: true,	
		},
		{
			name: "Register user with invalid phone number",
			args: args{
				name: "test",
				email: "test@mail.com",
				password: "test123",
				phoneNumber: "123456789",
			},
			wantErr: true,
		},
		{
			name: "Register user with short password",
			args: args{
				name: "test",
				email: "test@gmail.com",
				password: "test",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Reigster user with empty name",
			args: args{
				name: "",
				email: "tesr@mail.com",
				password: "test1234",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with empty email",
			args: args{
				name: "test",
				email: "",
				password: "test1234",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with empty password",
			args: args{
				name: "test",
				email: "test@mail.com",
				password: "",
				phoneNumber: "1234567890",
			},
			wantErr: true,
		},
		{
			name: "Register user with empty phone number",
			args: args{
				name: "test",
				email: "test@mail.com",
				password: "test1234",
				phoneNumber: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
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
	type args struct {
		email 	  		 string
		password  		 string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Login user with valid details",
			args: args{
				email: "test@gmail.com",
				password: "test1234",
			},
			wantErr: false,
		},
		{
			name: "Login user with invalid email",
			args: args{
				email: "test.user",
				password: "test123",
			},
			wantErr: true,
		},
		{
			name: "Login user with invalid password",
			args: args{
				email: "test@mail.com",
				password: "test1233",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			token, err := LoginUser(tt.args.email, tt.args.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NotEmpty(t, token)
				assert.NoError(t, err)
			}
		})
	}
}