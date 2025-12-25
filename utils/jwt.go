package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)
const secretKey="secretKey"
func GenerateJWT(userEmail string, userID int64) (string, error) {
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": userEmail,
		"userId":userID,
		"exp":     time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}