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

	// generate random id for state identification
	generated := utils.GenerateRandomString(32)

	sess, err := session.Get("session_state", c)
	if err != nil {
		log.Printf("Failed to initiate session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to initiate session")
	}

	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 10, // 10 minutes
		HttpOnly: true,
	}

	// set state value for verification in callback service
	sess.Values["state"] = generated
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Failed to initiate session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to initiate session")
	}

	URL := conf.AuthCodeURL(generated)

	return c.Redirect(http.StatusTemporaryRedirect, URL)
}

func GoogleLoginCallbackService(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")

	if code == "" || state == "" {
		return c.String(http.StatusUnauthorized, "Unauthorized")
	}

	stateSession, err := session.Get("session_state", c)
	if state != stateSession.Values["state"] {
		return c.String(http.StatusUnauthorized, "Session expired or invalid")
	}

	userSession, err := session.Get("session", c)
	if err != nil {
		log.Printf("Failed to initiate session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to initiate session")
	}

	userSession.Options = &sessions.Options{
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
	existingUser, _ := repos.FindUserByEmail(payload.Email)
	if existingUser == nil {
		newUser, err := CreateUser(payload)
		if err != nil {
			log.Printf("Failed to register user: %v", err)
			return c.String(http.StatusInternalServerError, "Failed to register user")
		}

		userSession.Values["ID"] = newUser.ID
		if err := userSession.Save(c.Request(), c.Response()); err != nil {
			log.Printf("Failed to initiate session: %v", err)
			return c.String(http.StatusInternalServerError, "Failed to initiate session")
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	userSession.Values["ID"] = existingUser.ID
	if err := userSession.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Failed to initiate session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to initiate session")
	}

	// Clean up state session as it's no longer needed.
	// Even though the lifetime of the stateSession is only 10 minutes,
	// cleaning up manually like this ensures there is no leak of state session, making it more secure.
	stateSession.Options.MaxAge = -1
	if err := stateSession.Save(c.Request(), c.Response()); err != nil {
		log.Printf("Failed to delete state session: %v", err)
		return c.String(http.StatusInternalServerError, "Failed to clean up session")
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}
