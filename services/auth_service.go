package services

import (
	"binder/configs"
	"binder/repos"
	"binder/utils"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func GoogleLoginService(c echo.Context) error {
	var conf = configs.GoogleConfig
	URL := conf.AuthCodeURL("not-implemented-yet")

	return c.Redirect(http.StatusTemporaryRedirect, URL)
}

func GoogleLoginCallbackService(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.NoContent(http.StatusUnauthorized)
	}

	sess, err := session.Get("session", c)
	if err != nil {
		log.Printf("Failed to initiate session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to initiate session")
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	// exchange token retrieved from google with valid token
	token, err := configs.GoogleConfig.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	// convert token to user profile data
	payload, err := utils.ConvertGoogleToken(token.AccessToken)
	if err != nil {
		return c.NoContent(http.StatusUnauthorized)
	}

	// check if there is a user recorded with the same creds
	exist, _ := repos.FindUserByEmail(payload.Email)
	if exist == nil {
		newUser, err := CreateUser(payload)
		if err != nil {
			log.Printf("Failed to register user: %v", err)
			return c.String(http.StatusInternalServerError, "Failed to register user")
		}

		sess.Values["ID"] = newUser.ID

		if err := sess.Save(c.Request(), c.Response()); err != nil {
			log.Printf("Failed to initiate session: %v", err)
			return c.String(http.StatusInternalServerError, "Failed to initiate session")
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	sess.Values["ID"] = exist.ID
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Failed to initiate session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to initiate session")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
