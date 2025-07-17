package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetServices godoc
// @Summary Get all services
// @Description Get all services
// @Tags services
// @Produce  json
// @Success 200 {array} models.Service
// @Router /services [get]
func GetServices(c *fiber.Ctx) error {
	services, err := repository.GetServices()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(services)
}

// GetService godoc
// @Summary Get a service by ID
// @Description Get a service by ID
// @Tags services
// @Produce  json
// @Param id path string true "Service ID"
// @Success 200 {object} models.Service
// @Router /services/{id} [get]
func GetService(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	service, err := repository.GetServiceByID(objID)
	if err != nil {
		return c.Status(404).SendString("Data tidak ditemukan")
	}
	return c.JSON(service)
}

// CreateService godoc
// @Summary Create a new service
// @Description Create a new service
// @Tags services
// @Accept  json
// @Produce  json
// @Param service body models.Service true "Service"
// @Success 201 {object} models.Service
// @Router /services [post]
func CreateService(c *fiber.Ctx) error {
	var service models.Service
	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Validasi 
	if service.Name == "" || service.Price < 0 {
		return c.Status(400).SendString("Nama atau harga tidak valid")
	}

	service.ID = primitive.NewObjectID()
	err := repository.CreateService(service)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(service)
}

// UpdateService godoc
// @Summary Update a service
// @Description Update a service
// @Tags services
// @Accept  json
// @Produce  json
// @Param id path string true "Service ID"
// @Param service body models.Service true "Service"
// @Success 200 {object} map[string]interface{}
// @Router /services/{id} [put]
func UpdateService(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}

	var service models.Service
	if err := c.BodyParser(&service); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	update := bson.M{"$set": bson.M{
		"name":        service.Name,
		"description": service.Description,
		"price":       service.Price,
	}}

	err = repository.UpdateService(objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

// DeleteService godoc
// @Summary Delete a service
// @Description Delete a service
// @Tags services
// @Produce  json
// @Param id path string true "Service ID"
// @Success 200 {object} map[string]interface{}
// @Router /services/{id} [delete]
func DeleteService(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	err = repository.DeleteService(objID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
}
