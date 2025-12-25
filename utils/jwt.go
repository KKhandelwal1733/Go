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

func VerifyToken(tokenString string) (int64,error) {
	parsedToken,err:=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {	
			return nil, jwt.ErrSignatureInvalid
		}	
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0,err
	}
	if !parsedToken.Valid {
		return 0,jwt.ErrTokenInvalidClaims
	}
	claims,ok:=parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0,jwt.ErrTokenInvalidClaims
	}
	
	userId:= int64(claims["userId"].(float64))
	return userId, nil
}