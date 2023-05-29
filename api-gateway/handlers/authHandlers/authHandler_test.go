package authHandlers

import (
	"api-gateway/dependencies"
	"api-gateway/domain"
	mocks "api-gateway/mocks/authmocks"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	proto "api-gateway/proto/auth"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	grpc *mocks.AuthServiceClient
}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}
func (suite *AuthHandlerTestSuite) SetupTest() {
	suite.grpc = &mocks.AuthServiceClient{}
}

func (suite *AuthHandlerTestSuite) TearDownTest() {
	suite.grpc.AssertExpectations(suite.T())
}

func (suite *AuthHandlerTestSuite) TestRegisterUser() {
	t := suite.T()
	t.Run("expect to return 201 when user registered successfully", func(t *testing.T) {
		// Arrange
		requestBody := domain.RegisterUserRequest{
			Name:        "test1",
			Email:       "test1@mail.com",
			Password:    "test@1234",
			PhoneNumber: "9876543987",
			Role:        "USER",
		}

		expectedRequest := proto.RegisterUserRequest{
			Name:        requestBody.Name,
			Email:       requestBody.Email,
			Password:    requestBody.Password,
			PhoneNumber: requestBody.PhoneNumber,
			Role:        domain.RoleMap[requestBody.Role],
		}

		expectedResponse := proto.RegisterUserResponse{
			StatusCode: http.StatusCreated,
			Message: "User registered successfully",
		}

		resp := domain.Message{
			Message: expectedResponse.Message,
		}

		exp, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("error while marshalling expected response: %v", err)
		}

		jsonRequest := `{"name" : "test1", "email" : "test1@mail.com", "password" : "test@1234", "phone_number" : "9876543987" , "role":"USER"}`
		req := httptest.NewRequest("POST", "/register", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act

		suite.grpc.On("RegisterUser", context.Background(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := &dependencies.Dependencies{
			AuthService: suite.grpc,
		}
		// Assert
		handler := RegisterUser(deps.AuthService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusCreated, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})

	t.Run("expect to return 400 when request body is invalid", func(t *testing.T) {
		// Arrange
		requestBody := domain.RegisterUserRequest{
			Name:        "test1",
			Email:       "testting@mailcom",
			Password:   "test@1234",
			PhoneNumber: "9876543987",
			Role:        "USER",
		}

		expectedRequest := proto.RegisterUserRequest{
			Name:        requestBody.Name,
			Email:       requestBody.Email,
			Password:    requestBody.Password,
			PhoneNumber: requestBody.PhoneNumber,
			Role:        domain.RoleMap[requestBody.Role],
		}

		expectedResponse := proto.RegisterUserResponse{
			StatusCode: http.StatusBadRequest,
			Message: "invalid request body",
		}

		resp := domain.Message{
			Message: expectedResponse.Message,
		}

		exp, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("error while marshalling expected response: %v", err)
		}

		jsonRequest := `{"name" : "test1", "email" : "testting@mailcom", "password" : "test@1234", "phone_number" : "9876543987" , "role":"USER"}`
		req := httptest.NewRequest("POST", "/register", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act

		suite.grpc.On("RegisterUser", context.Background(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := &dependencies.Dependencies{
			AuthService: suite.grpc,
		}
		// Assert
		handler := RegisterUser(deps.AuthService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusBadRequest, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})
}


func (suite *AuthHandlerTestSuite) TestLoginUser() {
	t := suite.T()
	t.Run("expect to return 200 when user logged in successfully", func(t *testing.T) {
		// Arrange
		requestBody := domain.LoginUserRequest{
			Email:    "test2@mailing.com",
			Password: "test@1234",
		}

		expectedRequest := proto.LoginUserRequest{
			Email:    requestBody.Email,
			Password: requestBody.Password,
		}

		expectedResponse := proto.LoginUserResponse{
			StatusCode: http.StatusOK,
			Message: "User logged in successfully",
		}

		resp := domain.LoginUserResponse{
			Message: expectedResponse.Message,
			Token: "",
		}

		exp, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("error while marshalling expected response: %v", err)
		}

		jsonRequest := `{"email" : "test2@mailing.com", "password" : "test@1234"}`
		req := httptest.NewRequest("POST", "/login", strings.NewReader(jsonRequest))
		res := httptest.NewRecorder()

		// Act

		suite.grpc.On("LoginUser", context.Background(), &expectedRequest).Return(&expectedResponse, nil).Once()
		deps := &dependencies.Dependencies{
			AuthService: suite.grpc,
		}
		// Assert
		handler := LoginUser(deps.AuthService)
		handler.ServeHTTP(res, req)
		assert.Equal(t, http.StatusOK, res.Code)
		assert.Equal(t, string(exp), res.Body.String())
	})
}