package routes

import (
	"petshop-backend/controllers"
	"petshop-backend/handler"
	"petshop-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	api.Post("/register", handler.Register)
	api.Post("/login", handler.Login)

	// Protected routes untuk Pets
	pets := api.Group("/pets", middleware.Protected())
	pets.Get("/", controllers.GetPets)
	pets.Get(":/id", controllers.GetPet)
	pets.Post("/", controllers.CreatePet)
	pets.Put(":/id", controllers.UpdatePet)
	pets.Delete(":/id", controllers.DeletePet)

	// Routes untuk Owners
	owners := api.Group("/owners")
	owners.Get("/", controllers.GetOwners)
	owners.Get(":/id", controllers.GetOwner)
	owners.Post("/", controllers.CreateOwner)
	owners.Put(":/id", controllers.UpdateOwner)
	owners.Delete(":/id", controllers.DeleteOwner)

	// Routes untuk Appointments
	appointments := api.Group("/appointments")
	appointments.Get("/", controllers.GetAppointments)
	appointments.Get(":/id", controllers.GetAppointmentWithDetails)
	appointments.Post("/", controllers.CreateAppointment)
	appointments.Put(":/id", controllers.UpdateAppointment)
	appointments.Delete(":/id", controllers.DeleteAppointment)

	// Routes untuk Services
	services := api.Group("/services")
	services.Get("/", controllers.GetServices)
	services.Get(":/id", controllers.GetService)
	services.Post("/", controllers.CreateService)
	services.Put(":/id", controllers.UpdateService)
	services.Delete(":/id", controllers.DeleteService)
}
