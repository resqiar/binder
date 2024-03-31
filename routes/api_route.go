package routes

import (
	"binder/services"

	"github.com/labstack/echo/v4"
)

func InitAPIRoutes(e *echo.Echo) {
	auth := e.Group("auth")

	auth.GET("/google", services.GoogleLoginService)
	auth.GET("/google/callback", services.GoogleLoginCallbackService)
}
