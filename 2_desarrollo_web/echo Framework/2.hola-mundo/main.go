package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Static("/js", "public/js")
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.File("/", "public/index.html")

	// Start server
	e.Logger.Fatal(e.Start(":80"))
}
