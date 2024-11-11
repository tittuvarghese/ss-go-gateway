package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Firstname string `json:"firstname" validate:"required,min=3,max=100"`
	Lastname  string `json:"lastname" validate:"required,min=3,max=100"`
	Username  string `json:"username" validate:"required,min=3,max=100"`
	Password  string `json:"password" validate:"required,min=4,max=100"`
	Type      string `json:"type" validate:"required,min=3,max=100"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=4,max=100"`
}

type ClaimsWrapper struct {
	Claims Claims `json:"claims"`
}

type Claims struct {
	Data AuthTokenPayload `json:"data"`
	jwt.RegisteredClaims
}

type AuthTokenPayload struct {
	ID        uuid.UUID `json:"userid"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Type      string    `json:"type"`
}
