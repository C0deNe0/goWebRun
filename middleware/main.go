package main

import "github.com/labstack/echo"

func main() {
	e := echo.New()

	//loading the routes
	// routes.InitRoutes(e)

	//starting the server
	e.Logger.Fatal(e.Start(":8080"))
}
