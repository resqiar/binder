package routes

import (
	"binder/middlewares"
	"binder/repos"
	"binder/utils"
	"binder/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitWebRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.HomePage())
	})

	e.GET("/dashboard", func(c echo.Context) error {
		userID := utils.GetUserIDFromContext(c.Request().Context())
		exts, _ := repos.GetAllExts(userID)

		return utils.Render(c, http.StatusOK, pages.IndexPage(exts))
	}, middlewares.UnprotectedMiddleware)

	e.GET("/login", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.LoginPage())
	})

	e.GET("/create", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.CreateExtPage())
	}, middlewares.ProtectedMiddleware)

	e.GET("/ext/:slug", func(c echo.Context) error {
		slug := c.Param("slug")
		userID := utils.GetUserIDFromContext(c.Request().Context())

		ext, err := repos.GetExt(userID, slug)
		if err != nil {
			return utils.Render(c, http.StatusOK, pages.NotFoundPage())
		}

		return utils.Render(c, http.StatusOK, pages.DetailExtPage(ext))
	}, middlewares.ProtectedMiddleware)
}
