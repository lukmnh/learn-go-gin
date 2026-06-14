package dto

import "time"

type UserCreateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,password"`
}

type UserCreateResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}
