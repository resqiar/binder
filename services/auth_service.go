package services

import (
	"binder/configs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GoogleLoginService(c echo.Context) error {
	var conf = configs.GoogleConfig
	URL := conf.AuthCodeURL("not-implemented-yet")

	return c.Redirect(http.StatusTemporaryRedirect, URL)
}
