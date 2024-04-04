package middlewares

import (
	"binder/repos"
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func ProtectedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			log.Println(err)
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		userID := sess.Values["ID"]
		if userID == nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		valid, err := repos.FindUserByID(userID.(string))
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		c.Set("userID", valid.ID)

		// save user id from session into context key value
		ctx := context.WithValue(c.Request().Context(), "userID", valid.ID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
