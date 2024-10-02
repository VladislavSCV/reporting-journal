package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	g := echo.New()
	g.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	g.Logger.Fatal(g.Start(":5001"))
}
