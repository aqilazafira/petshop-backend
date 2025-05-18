package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAppointments(c *fiber.Ctx) error {
	appointments, err := repository.GetAppointments()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(appointments)
}

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