package main

import (
	"net/http"

	
	"github.com/labstack/echo/v4"
)

//creating the cookies using the golang mux and storing the

func setcookieHandler(c echo.Context) error{
	cookie := new(http.Cookie)
	cookie.Name = "neo"
	cookie.Value = "yoo!!!!"
	cookie.Path = ""
	cookie.Secure = true
	cookie.HttpOnly =true

	c.SetCookie(cookie)

	return c.String(http.StatusOK,"cookie set")
}


func getcookieHandler(c echo.Context) error{
	cookei,err := c.Cookie("neo")
	if err!=nil{
		return c.String(http.StatusBadRequest,"unable to get the cookie")
	}

	return c.String(http.StatusOK,"got the cookie "+cookei.Value)
}

func main(){
	e := echo.New()

	e.GET("/setcookie",setcookieHandler)
	e.GET("/getcookie",getcookieHandler)

	e.Logger.Fatal(e.Start(":8000"))
}