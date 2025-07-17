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
)

// @title Petshop API
// @version 1.0
// @description REST API Petshop dengan Fiber & MongoDB
// @host localhost:3000
// @BasePath /api

func main() {
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "supersecret")
	}
	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())

	config.ConnectDB()
	routes.SetupRoutes(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // dok API

	log.Println("Server is running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
