package routes

import (
	"petshop-backend/config/middleware"
	"petshop-backend/controllers"
	"petshop-backend/handler"

	"github.com/gofiber/fiber/v2"
	storage "github.com/supabase-community/storage-go"
)

func SetupRoutes(app *fiber.App, supabaseClient *storage.Client) {
	middleware := middleware.Middlewares("admin")
	api := app.Group("/api")

	// Auth
	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	// Protected routes untuk Pets
	pets := api.Group("/pets")
	pets.Get("/", controllers.GetPets)
	pets.Get("/:id", controllers.GetPet)
	pets.Post("/", middleware, func(c *fiber.Ctx) error { return controllers.CreatePet(c, supabaseClient) })
	pets.Put("/:id", middleware, func(c *fiber.Ctx) error { return controllers.UpdatePet(c, supabaseClient) })
	pets.Delete("/:id", middleware, controllers.DeletePet)

	// Routes untuk Owners
	owners := api.Group("/owners")
	owners.Get("/", controllers.GetOwners)
	owners.Get("/:id", controllers.GetOwner)
	owners.Post("/", middleware, controllers.CreateOwner)
	owners.Put("/:id", middleware, controllers.UpdateOwner)
	owners.Delete("/:id", middleware, controllers.DeleteOwner)

	// Routes untuk Appointments
	appointments := api.Group("/appointments")
	appointments.Get("/", controllers.GetAppointments)
	appointments.Get("/:id", controllers.GetAppointmentWithDetails)
	appointments.Post("/", middleware, controllers.CreateAppointment)
	appointments.Put("/:id", middleware, controllers.UpdateAppointment)
	appointments.Delete("/:id", middleware, controllers.DeleteAppointment)

	// Routes untuk Services
	services := api.Group("/services")
	services.Get("/", controllers.GetServices)
	services.Get("/:id", controllers.GetService)
	services.Post("/", middleware, controllers.CreateService)
	services.Put("/:id", middleware, controllers.UpdateService)
	services.Delete("/:id", middleware, controllers.DeleteService)

	// Routes untuk Adoptions
	adoptions := api.Group("/adoptions")
	adoptions.Get("/", controllers.GetAdoptions)
	adoptions.Get("/my", controllers.GetMyAdoptions) // New endpoint for user's own adoptions
	adoptions.Get("/status", controllers.GetAdoptionsByStatus)
	adoptions.Get("/pet/:pet_id", controllers.GetAdoptionsByPetID)
	adoptions.Get("/:id", controllers.GetAdoption)
	adoptions.Post("/", controllers.CreateAdoption)
	adoptions.Put("/:id", middleware, controllers.UpdateAdoption)
	adoptions.Delete("/:id", middleware, controllers.DeleteAdoption)
}
