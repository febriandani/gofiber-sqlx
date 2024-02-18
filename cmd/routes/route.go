package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"gofiber-sqlx/handler"
)

// SetupRoutes sets up the routes for the application.
func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler, log *logrus.Logger) {
	app.Post("/users", userHandler.CreateUserHandler)
	app.Get("/users/:id", userHandler.GetUserByIDHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/users/:id", userHandler.DeleteUserHandler)
	app.Get("/users/:offset/:limit", userHandler.GetUsersHandler)
	app.Get("/", func(c *fiber.Ctx) error {
		response := map[string]string{
			"message": "Welcome to Dockerized app",
		}
		return c.JSON(response)
	})

}
