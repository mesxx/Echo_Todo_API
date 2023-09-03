package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}
