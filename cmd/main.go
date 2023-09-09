package main

import (
	"log"

	"gofiber-sqlx/cmd/routes"
	"gofiber-sqlx/handler"
	"gofiber-sqlx/repository"
	"gofiber-sqlx/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

const (
	dbDriver = "postgres"
	dbSource = "host=localhost port=5432 user=postgres password=junior34 dbname=postgres sslmode=disable"
	port     = ":8080" // Replace with the desired port number
)

func main() {
	// Connect to the database
	db, err := sqlx.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize the repository, service, and handler
	userRepo := repository.NewUsertRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app, userHandler)
	// Start the Fiber server
	log.Fatal(app.Listen(port))
}
