package domain

import (
	proto "api-gateway/proto/auth"
)

type RegisterUserRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
}

type Message struct {
	Message string `json:"message"`
}

var RoleMap = map[string]proto.Role{
	"USER":  proto.Role_USER,
	"ADMIN": proto.Role_ADMIN,
}
