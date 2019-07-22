package main

import (
	"level11api"
	//"level11infrastructure"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	
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
	// ตอนที่ TEST ใช้ 8082 เพื่อไม่ให้ชน
	//e.Logger.Fatal(e.Start(":8081"))
}
