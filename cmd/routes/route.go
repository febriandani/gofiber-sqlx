package routes

import (
	"github.com/gofiber/fiber/v2"

	"gofiber-sqlx/handler"
)

// SetupRoutes sets up the routes for the application.
func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
	app.Post("/users", userHandler.CreateUserHandler)
	app.Get("/users/:id", userHandler.GetUserByIDHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/users/:id", userHandler.DeleteUserHandler)
	app.Get("/users/:offset/:limit", userHandler.GetUsersHandler)
}
