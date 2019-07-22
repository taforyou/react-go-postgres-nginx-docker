package main

import (
	"level11api"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"level11infrastructure"
)

func main() {
	//level11infrastructure.ConnDB()
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//e.Static("/", "static")

	// Routes
	e.GET("/test", level11api.TestFetch)
	e.GET("/requestCheckId", level11api.RequestCheckId)
	

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
