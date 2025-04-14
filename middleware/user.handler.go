package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JWTCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	var input LoginInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTCustomClaims{
		Email: input.Email,
		Role:  input.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	tokenStr, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": " Failed to generate token"})
	}
	return c.JSON(http.StatusOK, echo.Map{"token": tokenStr})

}
