package routes

import (
	"binder/middlewares"
	"binder/utils"
	"binder/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitWebRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.IndexPage())
	})

	e.GET("/login", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.LoginPage())
	})

	e.GET("/create", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.CreateExtPage())
	}, middlewares.ProtectedMiddleware)
}
