package database

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

func CreateSessionStore() *session.Store {
	store := session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Minute * 1,
	})
	return store
}
