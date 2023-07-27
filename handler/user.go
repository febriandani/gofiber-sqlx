package handler

import (
	"log"
	"strconv"

	mt "gofiber-sqlx/model/user"
	utils "gofiber-sqlx/model/utils"
	ts "gofiber-sqlx/service"

	"github.com/gofiber/fiber/v2"
)

// UserHandler represents the HTTP handler for user-related operations.
type UserHandler struct {
	userService ts.User
}

// NewUserHandler creates a new UserHandler instance.
func NewUserHandler(userService ts.User) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUserHandler handles the "create user" HTTP request.
func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	user := new(mt.User)
	if err := c.BodyParser(user); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, "Invalid request payload"))
	}

	createdUser, err := h.userService.CreateUser(user.Name, user.Email)
	if err != nil {
		log.Println("Error creating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(fiber.StatusInternalServerError, "Failed to create user"))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse(fiber.StatusCreated, "User created successfully", createdUser))
}

// GetUserByIDHandler handles the "get user by ID" HTTP request.
func (h *UserHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, "Invalid user ID"))
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		log.Println("Error fetching user:", err)
		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse(fiber.StatusNotFound, "User not found"))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, "User retrieved successfully", user))
}

// UpdateUserHandler handles the "update user" HTTP request.
func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, "Invalid user ID"))
	}

	user := new(mt.User)
	err = c.BodyParser(user)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, "Invalid request payload"))
	}

	err = h.userService.UpdateUser(userID, user.Name, user.Email) 
	if err != nil {
		log.Println("Error updating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(fiber.StatusInternalServerError, "Failed to update user"))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, "User updated successfully", nil))
}

// DeleteUserHandler handles the "delete user" HTTP request.
func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, "Invalid user ID"))
	}

	if err := h.userService.DeleteUser(userID); err != nil {
		log.Println("Error deleting user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(fiber.StatusInternalServerError, "Failed to delete user"))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, "User deleted successfully", nil))
}