package main

import (
	"docking-in/internal"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	internal.RegisterRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
