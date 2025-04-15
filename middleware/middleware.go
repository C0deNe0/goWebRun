package main

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)


//this is a wropper function for the middleware 
func JWTMiddleware(requiredRole string) echo.MiddlewareFunc {

	//and here we are returning a middleware to the called router
	return func(next echo.HandlerFunc) echo.HandlerFunc {

		//and this middleware is returning a function 
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing Authorizaiton header"})
			}
			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.ParseWithClaims(tokenStr, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
				}
				return JWT_SECRET, nil
			})
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired Error"})
			}
			claims := token.Claims.(*JWTCustomClaims)
			if requiredRole != "" && claims.Role != requiredRole {
				return c.JSON(http.StatusForbidden, echo.Map{"error": "Forbidden: Insufficient role"})

			}
			c.Set("user", claims)
			return next(c)
		}
	}
}
