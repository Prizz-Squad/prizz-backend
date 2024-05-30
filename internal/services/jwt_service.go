package services

import (
	"github.com/EraldCaka/prizz-backend/internal/types"
	"github.com/EraldCaka/prizz-backend/util"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(util.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
func CreateJWTToken(tokenData types.JWTToken, secretKey string, expirationTime time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"id":       tokenData.ID,
		"username": tokenData.Username,
		"exp":      time.Now().Add(expirationTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	/*
		TODO: IMPLEMENT A LOGIC TO STORE THE JWT TOKEN INSIDE THE DATABASE (optional)
	*/
	return token.SignedString([]byte(secretKey))
}
