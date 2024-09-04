package routes

import (
	"binder/middlewares"
	"binder/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitAPIRoutes(e *echo.Echo) {
	auth := e.Group("auth")
	auth.GET("/google", services.GoogleLoginService)
	auth.GET("/google/callback", services.GoogleLoginCallbackService)
	auth.GET("/logout", services.LogoutService)

	ext := e.Group("api/ext", middlewares.ProtectedMiddleware)
	ext.POST("/create", services.CreateExtensionService, middleware.BodyLimit("6M"))
	ext.POST("/edit", services.EditExtensionService, middleware.BodyLimit("6M"))
	ext.DELETE("/delete/:slug", services.DeleteExtensionService)
	ext.DELETE("/image/:slug/:imageId", services.DeleteExtensionImageService)
	ext.GET("/search", services.SearchExtensionService)

	mobile := e.Group("mobile")
	mobile.POST("/register", services.MobileRegisterService)
}
