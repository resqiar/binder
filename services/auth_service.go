package services

import (
	"binder/configs"
	"binder/repos"
	"binder/utils"
	"log"
	"net/http"

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
	exist, err := repos.FindUserByEmail(payload.Email)
	if err != nil {
		newUser, err := CreateUser(payload)
		if err != nil {
			log.Printf("Failed to register user: %v", err)
			return c.String(http.StatusInternalServerError, "Failed to register user")
		}

		log.Println(newUser)
	}

	log.Println(exist)

	return c.JSON(http.StatusOK, payload)
}
