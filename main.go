package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	v1 := e.Group("v1/api")
	{
		v1.GET("/intro", GetIntro)
	}
	e.Logger.Fatal(e.Start(":8080"))
}

func GetIntro(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Ok",
		"data":    "Mini application to get along with docker",
	})
}
