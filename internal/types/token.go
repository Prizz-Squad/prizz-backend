package types

import "github.com/golang-jwt/jwt/v5"

type JWTToken struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Expire   string `json:"exp"`
	jwt.RegisteredClaims
}
