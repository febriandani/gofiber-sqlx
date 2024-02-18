package main

import (
	"log"

	"gofiber-sqlx/cmd/routes"
	"gofiber-sqlx/handler"
	"gofiber-sqlx/infra"
	"gofiber-sqlx/repository"
	"gofiber-sqlx/service"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	dbDriver = "postgres"
	dbSource = "host=localhost port=5432 user=postgres password=junior34 dbname=postgres sslmode=disable"
	port     = ":8080" // Replace with the desired port number
)

func main() {
	logger := infra.NewLogger()

	// Connect to the database
	db, err := sqlx.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Initialize the repository, service, and handler
	userRepo := repository.NewUsertRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService, logger)

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app, userHandler, logger)

	// Start the Fiber server
	logger.WithField("StartApp", "gofiber").Info("server listen to port ", port)
	logger.WithFields(logrus.Fields{
		"ExternalID":      "240200155667",
		"ResponseCode":    "RC-00",
		"ResponseMessage": "Transaction Success",
		"TransactionID":   "24556"}).Info("Exec Transaction RTGS")
	logger.WithFields(logrus.Fields{
		"ExternalID":      "240200154779",
		"ResponseCode":    "RC-99",
		"ResponseMessage": "Transaction Failed. Saldo tidak cukup",
		"TransactionID":   "24555"}).Info("Exec Transaction RTGS")
	logger.Fatal(app.Listen(port))
	logger.WithField("StartApp", "gofiber").Info("server listen to port ", port)

}
