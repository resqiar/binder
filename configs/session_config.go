package configs

import (
	"os"

	"github.com/gorilla/sessions"
)

var SessionStore sessions.Store

func InitiateSession() *sessions.Store {
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	return &SessionStore
}
