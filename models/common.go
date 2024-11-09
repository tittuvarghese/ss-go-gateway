package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Firstname string
	Lastname  string
	Username  string
	Password  string
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
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
}
