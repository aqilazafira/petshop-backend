package controllers

import (
	"petshop-backend/models"
	"petshop-backend/pkg/validator"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetOwners godoc
// @Summary Get all owners
// @Description Get all owners
// @Tags owners
// @Produce  json
// @Success 200 {array} models.Owner
// @Router /owners [get]
func GetOwners(c *fiber.Ctx) error {
	owners, err := repository.GetOwners()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(owners)
}

// GetOwner godoc
// @Summary Get an owner by ID
// @Description Get an owner by ID
// @Tags owners
// @Produce  json
// @Param id path string true "Owner ID"
// @Success 200 {object} models.Owner
// @Router /owners/{id} [get]
func GetOwner(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	owner, err := repository.GetOwnerByID(objID)
	if err != nil {
		return c.Status(404).SendString("Data tidak ditemukan")
	}
	return c.JSON(owner)
}

// CreateOwner godoc
// @Summary Create a new owner
// @Description Create a new owner
// @Tags owners
// @Accept  json
// @Produce  json
// @Param owner body models.Owner true "Owner"
// @Success 201 {object} models.Owner
// @Router /owners [post]
func CreateOwner(c *fiber.Ctx) error {
	var owner models.Owner
	if err := c.BodyParser(&owner); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Validate email format
	if isValid, errorMsg := validator.ValidateEmailFormat(owner.Email); !isValid {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid email format",
			"message": errorMsg,
		})
	}

	owner.ID = primitive.NewObjectID()
	err := repository.CreateOwner(owner)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(owner)
}

// UpdateOwner godoc
// @Summary Update an owner
// @Description Update an owner
// @Tags owners
// @Accept  json
// @Produce  json
// @Param id path string true "Owner ID"
// @Param owner body models.Owner true "Owner"
// @Success 200 {object} map[string]interface{}
// @Router /owners/{id} [put]
func UpdateOwner(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}

	var owner models.Owner
	if err := c.BodyParser(&owner); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	// Validate email format
	if isValid, errorMsg := validator.ValidateEmailFormat(owner.Email); !isValid {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid email format",
			"message": errorMsg,
		})
	}

	update := bson.M{"$set": bson.M{
		"name":    owner.Name,
		"email":   owner.Email,
		"phone":   owner.Phone,
		"address": owner.Address,
	}}

	err = repository.UpdateOwner(objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

// DeleteOwner godoc
// @Summary Delete an owner
// @Description Delete an owner
// @Tags owners
// @Produce  json
// @Param id path string true "Owner ID"
// @Success 200 {object} map[string]interface{}
// @Router /owners/{id} [delete]
func DeleteOwner(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	err = repository.DeleteOwner(objID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
}
