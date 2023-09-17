package controller

import (
	"github.com/biskitsx/Server-Side-Session/database"
	"github.com/biskitsx/Server-Side-Session/model"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
}
type userController struct{}

func NewUserController() UserController {
	return &userController{}
}

func (userController) GetUser(c *fiber.Ctx) error {
	users := []model.User{}
	database.Db.Find(&users)
	return c.JSON(users)
}
