package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAppointments godoc
// @Summary Get all appointments
// @Description Get all appointments
// @Tags appointments
// @Produce  json
// @Success 200 {array} models.Appointment
// @Router /appointments [get]
func GetAppointments(c *fiber.Ctx) error {
	appointments, err := repository.GetAppointments()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(appointments)
}

// GetAppointmentWithDetails godoc
// @Summary Get an appointment by ID with details
// @Description Get an appointment by ID with details
// @Tags appointments
// @Produce  json
// @Param id path string true "Appointment ID"
// @Success 200 {object} object
// @Router /appointments/{id} [get]
func GetAppointmentWithDetails(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}

	appointment, err := repository.GetAppointmentByID(objID)
	if err != nil {
		return c.Status(404).SendString("Data janji temu tidak ditemukan")
	}

	pet, _ := repository.GetPetByID(appointment.PetID)
	service, _ := repository.GetServiceByID(appointment.ServiceID)

	type AppointmentDetail struct {
		Appointment models.Appointment `json:"appointment"`
		Pet         models.Pet         `json:"pet"`
		Service     models.Service     `json:"service"`
	}

	detail := AppointmentDetail{
		Appointment: appointment,
		Pet:        pet,
		Service:    service,
	}

	return c.JSON(detail)
}

// CreateAppointment godoc
// @Summary Create a new appointment
// @Description Create a new appointment
// @Tags appointments
// @Accept  json
// @Produce  json
// @Param appointment body models.Appointment true "Appointment"
// @Success 201 {object} models.Appointment
// @Router /appointments [post]
func CreateAppointment(c *fiber.Ctx) error {
	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Validasi
	if appointment.PetID.IsZero() || appointment.ServiceID.IsZero() || appointment.Date == "" {
		return c.Status(400).SendString("Field wajib tidak boleh kosong")
	}

	appointment.ID = primitive.NewObjectID()
	err := repository.CreateAppointment(appointment)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(appointment)
}

// UpdateAppointment godoc
// @Summary Update an appointment
// @Description Update an appointment
// @Tags appointments
// @Accept  json
// @Produce  json
// @Param id path string true "Appointment ID"
// @Param appointment body models.Appointment true "Appointment"
// @Success 200 {object} map[string]interface{}
// @Router /appointments/{id} [put]
func UpdateAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}

	var appointment models.Appointment
	if err := c.BodyParser(&appointment); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	update := bson.M{"$set": bson.M{
		"pet_id":     appointment.PetID,
		"service_id": appointment.ServiceID,
		"date":       appointment.Date,
		"note":       appointment.Note,
	}}

	err = repository.UpdateAppointment(objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

// DeleteAppointment godoc
// @Summary Delete an appointment
// @Description Delete an appointment
// @Tags appointments
// @Produce  json
// @Param id path string true "Appointment ID"
// @Success 200 {object} map[string]interface{}
// @Router /appointments/{id} [delete]
func DeleteAppointment(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	err = repository.DeleteAppointment(objID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
}
