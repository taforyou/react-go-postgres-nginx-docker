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

	e.Static("/", "static")

	// Routes
	e.POST("/users", level11api.CreateUser)
	e.GET("/users", level11api.GetUsers)
	e.GET("/test", level11api.TestFetch)
	// การบ้านอูด้ง ไป Query คนเดียวมา
	// การบ้านอู้ง ไป Edit คนเดียวมา
	// ไปลบคนเดียวมา
	// ลบ เฉพาะ

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
