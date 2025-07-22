package routes

import (
	"petshop-backend/config/middleware"
	"petshop-backend/controllers"
	"petshop-backend/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Auth
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Protected routes untuk Pets
	pets := api.Group("/pets")
	pets.Get("/", controllers.GetPets)
	pets.Get(":/id", controllers.GetPet)
	pets.Post("/", controllers.CreatePet)
	pets.Put(":/id", controllers.UpdatePet)
	pets.Delete(":/id", middleware.Middlewares("admin"), controllers.DeletePet)

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

	// Routes untuk Adoptions
	adoptions := api.Group("/adoptions")
	adoptions.Get("/", controllers.GetAdoptions)
	adoptions.Get("/status", controllers.GetAdoptionsByStatus)
	adoptions.Get("/pet/:pet_id", controllers.GetAdoptionsByPetID)
	adoptions.Get("/:id", controllers.GetAdoption)
	adoptions.Post("/", controllers.CreateAdoption)
	adoptions.Put("/:id", controllers.UpdateAdoption)
	adoptions.Delete("/:id", controllers.DeleteAdoption)
}
