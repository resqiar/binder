package routes

import (
	"binder/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitAPIRoutes(e *echo.Echo) {
	auth := e.Group("auth")

	auth.GET("/google", services.GoogleLoginService)
	auth.GET("/google/callback", func(c echo.Context) error {
		return c.String(http.StatusOK, "Google Callback")
	})
}
