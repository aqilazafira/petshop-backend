package main

import (
	"log"
	"os"

	"petshop-backend/config"
	"petshop-backend/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	_ "petshop-backend/docs" 
)

// @title Petshop API
// @version 1.0
// @description This is a sample server Petshop server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api
// @schemes http
func main() {
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "supersecret")
	}
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()
	routes.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // serve API docs

	log.Println("Server is running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}