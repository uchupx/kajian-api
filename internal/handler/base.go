package handler

import "github.com/labstack/echo/v4"

type HandlerInterface interface {
	Routes(e *echo.Echo)
}
