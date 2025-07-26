package main

import (
	"log"
	"os"
	"strings"

	"petshop-backend/config"
	"petshop-backend/routes"

	_ "petshop-backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/gofiber/swagger"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			// Error loading .env file
		}
	}
}

// @title TES SWAGGER PEMROGRAMAN III
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/Fadhail
// @contact.email 71423044@std.ulbi.ac.id

// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// Connect to database
	config.ConnectDB()

	// Initialize Supabase client
	config.InitSupabase()

	app := fiber.New()

	// Logging request
	app.Use(logger.New())

	// Basic CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","), AllowCredentials: true,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // serve API docs

	// Setup router
	routes.SetupRoutes(app, config.SupabaseClient)

	// 404 handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Endpoint not found",
		})
	})

	// Baca PORT dari environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // default port kalau tidak ada
	}

	log.Printf("Server is running at http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
