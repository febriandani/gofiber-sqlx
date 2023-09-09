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
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid request payload",
			"id": "Muatan permintaan tidak valid"},
		))
	}

	createdUser, err := h.userService.CreateUser(user.Name, user.Email)
	if err != nil {
		log.Println("Error creating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(fiber.StatusInternalServerError, map[string]string{
			"en": "Failed to create user",
			"id": "Gagal membuat pengguna",
		},
		))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.SuccessResponse(fiber.StatusCreated, map[string]string{
		"en": "User created successfully",
		"id": "Pengguna berhasil dibuat",
	},
		createdUser))
}

// GetUserByIDHandler handles the "get user by ID" HTTP request.
func (h *UserHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid user ID",
			"id": "ID Pengguna tidak valid",
		}))
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		log.Println("Error fetching user:", err)
		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse(fiber.StatusNotFound, map[string]string{
			"en": "User not found",
			"id": "Pengguna tidak ditemukan",
		}))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, map[string]string{
		"en": "User retrieved successfully",
		"id": "Pengguna berhasil diambil",
	}, user))
}

// UpdateUserHandler handles the "update user" HTTP request.
func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid user ID",
			"id": "ID Pengguna tidak valid",
		}))
	}

	user := new(mt.User)
	err = c.BodyParser(user)
	if err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid request payload",
			"id": "Muatan permintaan tidak valid",
		}))
	}

	err = h.userService.UpdateUser(userID, user.Name, user.Email)
	if err != nil {
		log.Println("Error updating user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(fiber.StatusInternalServerError, map[string]string{
			"en": "Failed to update user",
			"id": "Gagal memperbarui pengguna",
		}))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, map[string]string{
		"en": "User updated successfully",
		"id": "Pengguna berhasil diperbarui",
	}, nil))
}

// DeleteUserHandler handles the "delete user" HTTP request.
func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error parsing user ID:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid user ID",
			"id": "ID Pengguna tidak valid",
		}))
	}

	if err := h.userService.DeleteUser(userID); err != nil {
		log.Println("Error deleting user:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse(fiber.StatusInternalServerError, map[string]string{
			"en": "Failed to delete user",
			"id": "Gagal menghapus pengguna",
		}))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, map[string]string{
		"en": "User deleted successfully",
		"id": "Pengguna berhasil dihapus",
	}, nil))
}

func (h *UserHandler) GetUsersHandler(c *fiber.Ctx) error {
	offset, err := strconv.Atoi(c.Params("offset"))
	if err != nil {
		log.Println("Error parsing offset:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid offset",
			"id": "Offset tidak valid",
		}))
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		log.Println("Error parsing limit:", err)
		return c.Status(fiber.StatusBadRequest).JSON(utils.ErrorResponse(fiber.StatusBadRequest, map[string]string{
			"en": "Invalid limit",
			"id": "Limit tidak valid",
		}))
	}

	user, err := h.userService.GetUsers(offset, limit)
	if err != nil {
		log.Println("Error fetching users:", err)
		return c.Status(fiber.StatusNotFound).JSON(utils.ErrorResponse(fiber.StatusNotFound, map[string]string{
			"en": "User not found",
			"id": "Pengguna tidak ditemukan",
		}))
	}

	return c.Status(fiber.StatusOK).JSON(utils.SuccessResponse(fiber.StatusOK, map[string]string{
		"en": "User retrieved successfully",
		"id": "Pengguna berhasil diambil",
	}, user))
}
