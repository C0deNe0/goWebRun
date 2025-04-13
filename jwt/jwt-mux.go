package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwtToken(userId string, userName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":    userId,
		"user_name":  userName,
		"iat":        time.Now().Unix(),
		"exp":        time.Now().Add(time.Hour * 24 * 360).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := os.Getenv("SECRET_KEY")

	if secretKey == "" {
		return "", fmt.Errorf("SECRET_KEY env not found")
	}

	signedToken, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func DecodeJwtToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secretKey := os.Getenv("SECRET_KEY")

		if secretKey == "" {
			return nil, fmt.Errorf("SECRET_KEY env not found")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")

}
