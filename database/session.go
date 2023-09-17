package database

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func CreateSessionStore() {
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Minute * 1,
	})
}
