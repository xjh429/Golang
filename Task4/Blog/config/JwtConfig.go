package config

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken 生成 Token
func GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(8 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// VerifyToken 验证 Token
func VerifyToken(tokenString string) (uint, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return 0, false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, false
	}
	if userIDVal, ok := claims["userID"]; ok {
		if userIDFloat, ok := userIDVal.(float64); ok {
			return uint(userIDFloat), true
		}
	}
	return 0, false
}
