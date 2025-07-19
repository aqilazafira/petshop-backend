package handler

import (
	"fmt"
	"petshop-backend/config/middleware"
	"petshop-backend/models"
	pwd "petshop-backend/pkg/password"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var req models.UserLogin

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	user, err := repository.FindUserByEmail(c.Context(), req.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Email not found"})
	}

	// Cek password input hash yang tersimpan
	if !pwd.CheckPasswordHash(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Wrong password"})
	}

	// Generate token PASETO
	token, err := middleware.EncodeWithRoleHours(user.Role, user.Username, 2)
	if err != nil {
		fmt.Println("Token generation error:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
	})
}

func Register(c *fiber.Ctx) error {
	var req models.UserLogin

	// Debug: Print headers and body
	fmt.Printf("Content-Type: %s\n", c.Get("Content-Type"))
	fmt.Printf("Raw body: %s\n", string(c.Body()))

	if err := c.BodyParser(&req); err != nil {
		fmt.Printf("Body parsing error: %v\n", err)
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   "Invalid JSON format",
			"details": err.Error(),
		})
	}

	if req.Email == "" || req.Username == "" || req.Password == "" || req.Role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email, username, password, and role are required"})
	}

	hashed, err := pwd.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	req.Password = hashed

	id, err := repository.InsertUser(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"id":      id,
	})
}
