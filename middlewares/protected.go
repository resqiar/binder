package middlewares

import (
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
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		userID := sess.Values["ID"]
		if userID == nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		// save user id from session into local key value
		c.Set("userID", userID)

		return next(c)
	}
}
