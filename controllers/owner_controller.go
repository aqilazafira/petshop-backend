package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOwners(c *fiber.Ctx) error {
	owners, err := repository.GetOwners()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(owners)
}

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

func CreateOwner(c *fiber.Ctx) error {
	var owner models.Owner
	if err := c.BodyParser(&owner); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	owner.ID = primitive.NewObjectID()
	err := repository.CreateOwner(owner)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.Status(201).JSON(owner)
}

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

	update := bson.M{"$set": bson.M{
		"name":  owner.Name,
		"email": owner.Email,
		"phone": owner.Phone,
	}}

	err = repository.UpdateOwner(objID, update)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

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
