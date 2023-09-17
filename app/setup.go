package app

import (
	"log"

	"github.com/biskitsx/Server-Side-Session/container"
	"github.com/biskitsx/Server-Side-Session/controller"
	"github.com/biskitsx/Server-Side-Session/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func SetupAndRunApp() error {
	// dotenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	db := database.ConnectPostgres()
	session := database.CreateSessionStore()
	container := container.NewContainer(db, session)

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "Access-Control-Allow-Origin, Content-Type, Origin, Accept",
	}))

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello")
	})

	authController := controller.NewAuthController(container)
	app.Post("/login", authController.Login)
	app.Post("/register", authController.Register)
	app.Post("/logout", authController.Logout)
	app.Get("/healthcheck", authController.HealthCheck)

	userController := controller.NewUserController(container)
	app.Get("/user", userController.GetUser)
	app.Listen(":8080")

	return nil
}
