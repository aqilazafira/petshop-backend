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
		return c.Status(400).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Validation
	if adoption.Name == "" || adoption.Email == "" || adoption.Phone == "" ||
		adoption.Address == "" || adoption.Reason == "" || adoption.LivingSpace == "" {
		return c.Status(400).JSON(fiber.Map{
			"error":   "Missing required fields",
			"message": "Name, email, phone, address, reason, and living space are required",
		})
	}

	// Set default values
	adoption.ID = primitive.NewObjectID()
	adoption.Status = "pending"
	adoption.SubmissionDate = time.Now()
	adoption.CreatedAt = time.Now()
	adoption.UpdatedAt = time.Now()

	err := repository.CreateAdoption(adoption)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to create adoption",
			"details": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Adoption request submitted successfully",
		"data":    adoption,
	})
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
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}

	var adoption models.Adoption
	if err := c.BodyParser(&adoption); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	update := bson.M{"$set": bson.M{
		"status":     adoption.Status,
		"updated_at": time.Now(),
	}}

	// If adoption is approved, set adoption date
	if adoption.Status == "approved" {
		now := time.Now()
		update["$set"].(bson.M)["adoption_date"] = now
	}

	err = repository.UpdateAdoption(objID, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update adoption",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Adoption updated successfully",
	})
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
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid ID format",
		})
	}
	err = repository.DeleteAdoption(objID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete adoption",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Adoption deleted successfully",
	})
}

// GetAdoptionsByStatus godoc
// @Summary Get adoptions by status
// @Description Get all adoptions filtered by status
// @Tags adoptions
// @Produce  json
// @Param status query string true "Status (pending, approved, rejected)"
// @Success 200 {array} models.Adoption
// @Router /adoptions/status [get]
func GetAdoptionsByStatus(c *fiber.Ctx) error {
	status := c.Query("status")
	if status == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "Status parameter is required",
		})
	}

	adoptions, err := repository.GetAdoptionsByStatus(status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch adoptions",
		})
	}

	return c.JSON(fiber.Map{
		"data":  adoptions,
		"count": len(adoptions),
	})
}

// GetAdoptionsByPetID godoc
// @Summary Get adoptions by pet ID
// @Description Get all adoptions for a specific pet
// @Tags adoptions
// @Produce  json
// @Param pet_id path string true "Pet ID"
// @Success 200 {array} models.Adoption
// @Router /adoptions/pet/{pet_id} [get]
func GetAdoptionsByPetID(c *fiber.Ctx) error {
	petIDStr := c.Params("pet_id")
	petID, err := primitive.ObjectIDFromHex(petIDStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid pet ID format",
		})
	}

	adoptions, err := repository.GetAdoptionsByPetID(petID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to fetch adoptions",
		})
	}

	return c.JSON(fiber.Map{
		"data":  adoptions,
		"count": len(adoptions),
	})
}
