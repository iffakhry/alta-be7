package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/hello", HelloHandler)
	e.Start(":8080")
}

func HelloHandler(c echo.Context) error {
	response := map[string]string{
		"data": "hello world",
	}
	return c.JSON(http.StatusOK, response)
}
