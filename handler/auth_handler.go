package handler

import (
	"petshop-backend/config/middleware"
	"petshop-backend/models"
	pwd "petshop-backend/pkg/password"
	"petshop-backend/repository"

	"github.com/gofiber/fiber/v2"
)

// Login godoc
// @Summary User login
// @Description User login with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body models.UserLogin true "Login details"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /login [post]
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{
		"message": "Login success",
		"token":   token,
		"user": fiber.Map{
			"email":    user.Email,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

// Register godoc
// @Summary User registration
// @Description User registration with email, username, password, and role
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body models.UserLogin true "User registration details"
// @Success 201 {object} map[string]interface{} "User registered successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 409 {object} map[string]interface{} "Conflict"
// @Failure 422 {object} map[string]interface{} "Unprocessable Entity"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /register [post]
func Register(c *fiber.Ctx) error {
	var req models.UserLogin

	if err := c.BodyParser(&req); err != nil {
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
