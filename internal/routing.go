package internal

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/kajian-api/internal/handler"
)

func InitRoute(e *echo.Echo) {
	internal := Internal{}
	handlers := []handler.HandlerInterface{
		internal.GetKajianHandler(),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	for _, h := range handlers {
		h.Routes(e)
	}
}
