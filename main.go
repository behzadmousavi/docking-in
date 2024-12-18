package main

import (
	"docking-in/internal/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	server.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
