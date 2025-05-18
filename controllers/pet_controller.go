package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPets(c *fiber.Ctx) error {
	pets, err := repository.GetPets()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(pets)
}

func GetPet(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	pet, err := repository.GetPetByID(objID)
	if err != nil {
		return c.Status(404).SendString("Data tidak ditemukan")
	}
	return c.JSON(pet)
}

func CreatePet(c *fiber.Ctx) error {
	var pet models.Pet
	if err := c.BodyParser(&pet); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	pet.ID = primitive.NewObjectID()
	err := repository.CreatePet(pet)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(pet)
}

func UpdatePet(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}

	var pet models.Pet
	if err := c.BodyParser(&pet); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	update := bson.M{"$set": bson.M{
		"name":     pet.Name,
		"species":  pet.Species,
		"age":      pet.Age,
		"owner_id": pet.OwnerID,
	}}

	err = repository.UpdatePet(objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

func DeletePet(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).SendString("ID tidak valid")
	}
	err = repository.DeletePet(objID)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(fiber.Map{"message": "Data berhasil dihapus"})
}