package main

import (
	"binder/utils"
	"binder/views/pages"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "views/public")

	e.GET("/", func(c echo.Context) error {
		return utils.Render(c, http.StatusOK, pages.IndexPage())
	})

	e.Logger.Fatal(e.Start(":8888"))
}
