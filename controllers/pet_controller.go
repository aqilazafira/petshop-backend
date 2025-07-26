package controllers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"petshop-backend/models"
	"petshop-backend/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	storage "github.com/supabase-community/storage-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// uploadToSupabaseStorage uploads a file to Supabase Storage using HTTP API
func uploadToSupabaseStorage(bucketName, fileName string, fileData []byte) (string, error) {
	// Get environment variables
	supabaseURL := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || apiKey == "" {
		return "", fmt.Errorf("SUPABASE_URL dan SUPABASE_KEY environment variables harus diset")
	}

	uploadURL := fmt.Sprintf("%s/storage/v1/object/%s/%s", supabaseURL, bucketName, fileName)

	// Create a buffer to write our multipart form data to
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Create a form file field
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}

	// Copy file data to the form file field
	_, err = io.Copy(part, bytes.NewReader(fileData))
	if err != nil {
		return "", err
	}

	// Close the writer to finalize the form data
	err = writer.Close()
	if err != nil {
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", uploadURL, &body)
	if err != nil {
		return "", err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("apikey", apiKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Return the public URL
	publicURL := fmt.Sprintf("%s/storage/v1/object/public/%s/%s", supabaseURL, bucketName, fileName)
	return publicURL, nil
}

// GetPets godoc
// @Summary Get all pets
// @Description Get all pets
// @Tags pets
// @Produce  json
// @Success 200 {array} models.Pet
// @Router /pets [get]
func GetPets(c *fiber.Ctx) error {
	pets, err := repository.GetPets()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(pets)
}

// GetPet godoc
// @Summary Get a pet by ID
// @Description Get a pet by ID
// @Tags pets
// @Produce  json
// @Param id path string true "Pet ID"
// @Success 200 {object} models.Pet
// @Router /pets/{id} [get]
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

// CreatePet godoc
// @Summary Create a new pet
// @Description Create a new pet
// @Tags pets
// @Accept  multipart/form-data
// @Produce  json
// @Param name formData string true "Pet Name"
// @Param species formData string true "Pet Species"
// @Param age formData integer true "Pet Age"
// @Param gender formData string true "Pet Gender"
// @Param owner_id formData string true "Owner ID"
// @Param image formData file false "Pet Image"
// @Success 201 {object} models.Pet
// @Router /pets [post]
func CreatePet(c *fiber.Ctx, supabaseClient *storage.Client) error {
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
		Status:  "available", // Set default status as available
	}

	file, err := c.FormFile("image")
	if err == nil {
		// Read the file content
		fileContent, err := file.Open()
		if err != nil {
			return c.Status(500).SendString("Gagal membuka file")
		}
		defer fileContent.Close()

		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			return c.Status(500).SendString("Gagal membaca file")
		}

		// Generate a unique filename
		fileName := uuid.New().String() + filepath.Ext(file.Filename)

		// Upload to Supabase Storage using custom HTTP function
		publicURL, err := uploadToSupabaseStorage(
			"pets",
			fileName,
			fileBytes,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   "Gagal mengunggah gambar ke Supabase",
				"details": err.Error(),
			})
		}

		pet.ImageURL = publicURL
	}

	if err := repository.CreatePet(pet); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(201).JSON(pet)
}

// UpdatePet godoc
// @Summary Update a pet
// @Description Update a pet
// @Tags pets
// @Accept  multipart/form-data
// @Produce  json
// @Param id path string true "Pet ID"
// @Param name formData string true "Pet Name"
// @Param species formData string true "Pet Species"
// @Param age formData integer true "Pet Age"
// @Param gender formData string true "Pet Gender"
// @Param owner_id formData string true "Owner ID"
// @Param image formData file false "Pet Image"
// @Success 200 {object} map[string]interface{}
// @Router /pets/{id} [put]
func UpdatePet(c *fiber.Ctx, supabaseClient *storage.Client) error {
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
		"name":    c.FormValue("name"),
		"species": c.FormValue("species"),
		"age":     age,
		"gender":  c.FormValue("gender"), "owner_id": ownerID,
	}

	file, err := c.FormFile("image")
	if err == nil {
		// Read the file content
		fileContent, err := file.Open()
		if err != nil {
			return c.Status(500).SendString("Gagal membuka file")
		}
		defer fileContent.Close()

		fileBytes, err := io.ReadAll(fileContent)
		if err != nil {
			return c.Status(500).SendString("Gagal membaca file")
		}

		// Generate a unique filename
		fileName := uuid.New().String() + filepath.Ext(file.Filename)

		// Upload to Supabase Storage using custom HTTP function
		publicURL, err := uploadToSupabaseStorage(
			"pets",
			fileName,
			fileBytes,
		)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error":   "Gagal mengunggah gambar ke Supabase",
				"details": err.Error(),
			})
		}

		update["image_url"] = publicURL
	}

	if err := repository.UpdatePet(objID, bson.M{"$set": update}); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(fiber.Map{"message": "Data berhasil diperbarui"})
}

// DeletePet godoc
// @Summary Delete a pet
// @Description Delete a pet
// @Tags pets
// @Produce  json
// @Param id path string true "Pet ID"
// @Success 200 {object} map[string]interface{}
// @Router /pets/{id} [delete]
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
