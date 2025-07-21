package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAdoptions godoc
// @Summary Get all adoptions
// @Description Get all adoptions
// @Tags adoptions
// @Produce  json
// @Success 200 {array} models.Adoption
// @Router /adoptions [get]
func GetAdoptions(c *fiber.Ctx) error {
	adoptions, err := repository.GetAdoptions()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(adoptions)
}

// GetAdoption godoc
// @Summary Get an adoption by ID
// @Description Get an adoption by ID
// @Tags adoptions
// @Produce  json
// @Param id path string true "Adoption ID"
// @Success 200 {object} models.Adoption
// @Router /adoptions/{id} [get]
func GetAdoption(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	adoption, err := repository.GetAdoptionByID(objID)
	if err != nil {
		return c.Status(404).SendString("Data tidak ditemukan")
	}
	return c.JSON(adoption)
}

// CreateAdoption godoc
// @Summary Create a new adoption
// @Description Create a new adoption
// @Tags adoptions
// @Accept  json
// @Produce  json
// @Param adoption body models.Adoption true "Adoption"
// @Success 201 {object} models.Adoption
// @Router /adoptions [post]
func CreateAdoption(c *fiber.Ctx) error {
	var adoption models.Adoption
	if err := c.BodyParser(&adoption); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	adoption.ID = primitive.NewObjectID()
	adoption.AdoptionDate = primitive.NewDateTimeFromTime(time.Now())
	err := repository.CreateAdoption(adoption)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(adoption)
}

// UpdateAdoption godoc
// @Summary Update an adoption
// @Description Update an adoption
// @Tags adoptions
// @Accept  json
// @Produce  json
// @Param id path string true "Adoption ID"
// @Param adoption body models.Adoption true "Adoption"
// @Success 200 {object} map[string]interface{}
// @Router /adoptions/{id} [put]
func UpdateAdoption(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}

	var adoption models.Adoption
	if err := c.BodyParser(&adoption); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	update := bson.M{"$set": bson.M{
		"pet_id":   adoption.PetID,
		"owner_id": adoption.OwnerID,
		"status":   adoption.Status,
	}}

	err = repository.UpdateAdoption(objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

// DeleteAdoption godoc
// @Summary Delete an adoption
// @Description Delete an adoption
// @Tags adoptions
// @Produce  json
// @Param id path string true "Adoption ID"
// @Success 200 {object} map[string]interface{}
// @Router /adoptions/{id} [delete]
func DeleteAdoption(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	err = repository.DeleteAdoption(objID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
}
