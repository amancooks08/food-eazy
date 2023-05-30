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

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type AddItemRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32   `json:"quantity"`
}

type AddItemResponse struct {
	ID          int32  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32   `json:"quantity"`
}

type GetItemRequest struct {
	ID int32 `json:"id"`
}

type GetItemResponse struct {
	ID          int32  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    uint32   `json:"quantity"`
}

type GetAllItemsResponse struct {
	Items []GetItemResponse `json:"items"`
}