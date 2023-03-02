package entity

import "time"

type CreateUserRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
}

type CreateUserResponse struct {
	Id          int64     `json:"id"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token    string    `json:"token"`
	ExpireAt time.Time `json:"expire_at"`
}
