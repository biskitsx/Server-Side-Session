package container

import (
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
)

type Container interface {
	GetDatabase() *gorm.DB
	GetStore() *session.Store
}
type container struct {
	db      *gorm.DB
	session *session.Store
}

func NewContainer(db *gorm.DB, session *session.Store) Container {
	return &container{
		db:      db,
		session: session,
	}
}

func (c *container) GetDatabase() *gorm.DB {
	return c.db
}

func (c *container) GetStore() *session.Store {
	return c.session
}
