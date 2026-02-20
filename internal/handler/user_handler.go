package handler

import (
	"blog_api/internal/model"
	"blog_api/internal/repository"
	"blog_api/internal/service"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service *service.AuthService
	Repo    *repository.UserRepository
}

func NewUserHandler(service *service.AuthService, repo *repository.UserRepository) *UserHandler {
	return &UserHandler{
		Service: service,
		Repo:    repo,
	}
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	role := c.Query("role")

	var users []model.User
	var err error

	if role != "" {
		users, err = h.Repo.FindByRole(role)
	} else {
		users, err = h.Repo.FindAll()
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data user"})
	}

	return c.JSON(fiber.Map{"data": users})
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userIDInterface := c.Locals("user_id")
	
	if userIDInterface == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Akses ditolak, silakan login kembali"})
	}

	userID := uint(userIDInterface.(float64))

	user, err := h.Service.GetProfile(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Berhasil mengambil data profil",
		"data":    user,
	})
}