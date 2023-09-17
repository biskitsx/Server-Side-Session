package controller

import (
	"github.com/biskitsx/Server-Side-Session/container"
	"github.com/biskitsx/Server-Side-Session/model"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	GetUser(c *fiber.Ctx) error
}
type userController struct {
	container container.Container
}

func NewUserController(c container.Container) UserController {
	return &userController{
		container: c,
	}
}

func (controller userController) GetUser(c *fiber.Ctx) error {
	users := []model.User{}
	db := controller.container.GetDatabase()
	db.Find(&users)
	return c.JSON(users)
}
