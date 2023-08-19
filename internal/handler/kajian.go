package handler

import "github.com/labstack/echo/v4"

type KajianHandler struct{}

// Routes is a function to register handler routes
func (h *KajianHandler) Routes(e *echo.Echo) {
	e.GET("/kajians", h.GetKajian)
}

func (h *KajianHandler) GetKajian(c echo.Context) error {
	return c.String(200, "Hello, World!")
}
