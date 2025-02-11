package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	var e = echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello, World!",
		})
	})

	if err := e.Start(":8000"); err != nil {
		log.Fatal(err)
	}

}
