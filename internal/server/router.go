package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetIntro(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  http.StatusOK,
		"message": "Ok",
		"data":    "Mini application to get along with docker",
	})
}

func RegisterRoutes(router *echo.Echo) {
	v1 := router.Group("v1/api")
	{
		v1.GET("/intro", GetIntro)
	}
}
