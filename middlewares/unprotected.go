package middlewares

import (
	"context"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func UnprotectedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		userID := sess.Values["ID"]

		// save user id from session into context key value
		ctx := context.WithValue(c.Request().Context(), "userID", userID)
		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
