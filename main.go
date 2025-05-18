package main

import (
	"log"

    "petshop-backend/config"
    "petshop-backend/routes"

    "github.com/gofiber/fiber/v2"
	 "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
    app := fiber.New()

	app.Use(cors.New())
    app.Use(logger.New())

    config.ConnectDB()
    routes.SetupRoutes(app)

    log.Println("Server is running on http://localhost:3000")
    if err := app.Listen(":3000"); err != nil {
        log.Fatalf("Gagal menjalankan server: %v", err)
    }
}