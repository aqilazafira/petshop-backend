package controllers

import (
	"petshop-backend/models"
	"petshop-backend/repository"
	"strconv"

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
	age, err := strconv.Atoi(c.FormValue("age"))
	if err != nil {
		return c.Status(400).SendString("Umur tidak valid")
	}

	ownerID, err := primitive.ObjectIDFromHex(c.FormValue("owner_id"))
	if err != nil {
		return c.Status(400).SendString("ID pemilik tidak valid")
	}

	pet := models.Pet{
		ID:      primitive.NewObjectID(),
		Name:    c.FormValue("name"),
		Species: c.FormValue("species"),
		Age:     age,
		Gender:  c.FormValue("gender"),
		OwnerID: ownerID,
	}

	file, err := c.FormFile("image")
	if err == nil {
		filePath := "../Petshop-App/public/pets/" + file.Filename
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(500).SendString("Gagal menyimpan file")
		}
		pet.ImageURL = "/pets/" + file.Filename
	}

	if err := repository.CreatePet(pet); err != nil {
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

	age, err := strconv.Atoi(c.FormValue("age"))
	if err != nil {
		return c.Status(400).SendString("Umur tidak valid")
	}

	ownerID, err := primitive.ObjectIDFromHex(c.FormValue("owner_id"))
	if err != nil {
		return c.Status(400).SendString("ID pemilik tidak valid")
	}

	update := bson.M{
		"name":     c.FormValue("name"),
		"species":  c.FormValue("species"),
		"age":      age,
		"gender":   c.FormValue("gender"),
		"owner_id": ownerID,
	}

	file, err := c.FormFile("image")
	if err == nil {
		filePath := "../Petshop-App/public/pets/" + file.Filename
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(500).SendString("Gagal menyimpan file")
		}
		update["image_url"] = "/pets/" + file.Filename
	}

	if err := repository.UpdatePet(objID, bson.M{"$set": update}); err != nil {
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
