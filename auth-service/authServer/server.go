package authServer

import (
	"context"
	"net/http"

	"auth-service/errors"
	proto "auth-service/proto/authpb"
	"auth-service/service"
)

type GRPCServer struct {
	proto.UnimplementedAuthServiceServer
}

func (s *GRPCServer) RegisterUser(ctx context.Context, req *proto.RegisterUserRequest) (*proto.RegisterUserResponse, error) {
	err := service.RegisterUser(req.Name, req.Email, req.Password, req.PhoneNumber, req.Role.String())
	if err != nil {
		return &proto.RegisterUserResponse{
			StatusCode: http.StatusBadRequest,
			Message:    "user not registered",
		}, err
	}

	return &proto.RegisterUserResponse{
		StatusCode: http.StatusOK,
		Message:    "User registered successfully",
	}, nil
}

func (s *GRPCServer) LoginUser(ctx context.Context, req *proto.LoginUserRequest) (*proto.LoginUserResponse, error) {
	token, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		if err.Error() == errors.ErrDuplicateEmail.Error() {
			return &proto.LoginUserResponse{
				StatusCode: http.StatusConflict,
				Token:      "",
			}, err
		}
	}

	return &proto.LoginUserResponse{
		StatusCode: http.StatusOK,
		Token:      token,
	}, nil
}

func (s *GRPCServer) ValidateUser(ctx context.Context, req *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error) {
	status, err := service.ValidateUser(req.Token, req.Role.String())
	if err != nil {
		return nil, err
	}

	return &proto.ValidateTokenResponse{
		StatusCode: int32(status),
		Message:    "Token is valid",
	}, nil
}
