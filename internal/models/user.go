package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type CreateUser struct {
	AuthUser
	User
}

type User struct {
	ID                uuid.UUID
	Role              string
	Name              string
	Surname           string
	Phone             string
	Address           string
	AddressCoordinate Coordinate
}

type AuthUser struct {
	Login    string
	Password string
}

type Coordinate struct {
	X float32
	Y float32
}

type UserInfo struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

type CustomClaims struct {
	jwt.StandardClaims
	UserInfo UserInfo `json:"user_info"`
}
