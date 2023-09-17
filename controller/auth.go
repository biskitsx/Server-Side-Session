package controller

import (
	"github.com/biskitsx/Server-Side-Session/database"
	"github.com/biskitsx/Server-Side-Session/model"
	"github.com/biskitsx/Server-Side-Session/service"
	"github.com/gofiber/fiber/v2"
)

type AuthController interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	HealthCheck(c *fiber.Ctx) error
}

type authController struct {
	authService service.AuthService
}

func NewAuthController() AuthController {
	return &authController{
		authService: service.NewAuthService(),
	}
}

type UserRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (controller *authController) Register(c *fiber.Ctx) error {
	userRegister := UserRegister{}

	if err := c.BodyParser(&userRegister); err != nil {
		return fiber.NewError(401, err.Error())
	}

	_, err := controller.authService.CheckUsername(userRegister.Username)
	if err == nil {
		return fiber.NewError(401, "This username already registered")
	}

	user := model.User{
		Username: userRegister.Username,
		Password: userRegister.Password,
	}

	database.Db.Create(&user)
	return c.JSON(user)
}

func (controller authController) Login(c *fiber.Ctx) error {
	userLogin := UserLogin{}

	if err := c.BodyParser(&userLogin); err != nil {
		return fiber.NewError(401, err.Error())
	}

	oldUser, err := controller.authService.CheckUsername(userLogin.Username)
	if err != nil {
		return fiber.NewError(401, "invalid username or password")
	}

	if oldUser.Password != userLogin.Password {
		return fiber.NewError(401, "invalid password")
	}

	sess, err := database.Store.Get(c)
	if err != nil {
		return fiber.NewError(401, "invalid password")
	}

	sess.Set("authenticated", true)
	sess.Set("user_id", oldUser.ID)

	if err := sess.Save(); err != nil {
		return fiber.NewError(401, "error while save session")
	}

	return c.JSON(fiber.Map{
		"msg": "logged in",
	})
}
func (controller authController) Logout(c *fiber.Ctx) error {
	sess, err := database.Store.Get(c)
	if err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "logged out (no session)",
		})
	}

	err = sess.Destroy()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "logged out",
	})
}

func (controller authController) HealthCheck(c *fiber.Ctx) error {
	sess, err := database.Store.Get(c)
	if err != nil {
		return fiber.NewError(401, "unauthorized ")
	}
	auth := sess.Get("authenticated")
	if auth == nil {
		return fiber.NewError(401, "unauthorized ")
	}
	userId := sess.Get("user_id")
	return c.JSON(fiber.Map{
		"auth":   auth,
		"userId": userId,
	})
}
