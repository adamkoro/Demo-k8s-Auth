package model

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
