package handler

import (
	"blog_api/internal/model"
	"blog_api/internal/repository"
	"blog_api/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostHandler struct {
	Service *service.PostService
	Repo    *repository.PostRepository
}

func NewPostHandler(service *service.PostService, repo *repository.PostRepository) *PostHandler {
	return &PostHandler{Service: service, Repo: repo}
}

func (h *PostHandler) CreatePost(c *fiber.Ctx) error {
	var input model.Post
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	if err := h.Service.CreatePost(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	fullPost, err := h.Repo.FindBySlug(input.Slug)
	if err != nil {
		return c.Status(201).JSON(fiber.Map{"message": "Artikel berhasil dibuat", "data": input})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Artikel berhasil dibuat", 
		"data": fullPost,
	})
}

func (h *PostHandler) UpdatePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID artikel tidak valid"})
	}

	var input model.Post
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input tidak valid"})
	}

	userID := input.UserID 

	updatedPost, err := h.Service.UpdatePost(uint(postID), userID, input)
	if err != nil {
		if err.Error() == "akses ditolak: kamu hanya bisa mengedit artikelmu sendiri" {
			return c.Status(403).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Artikel berhasil diperbarui",
		"data":    updatedPost,
	})
}

func (h *PostHandler) DeletePost(c *fiber.Ctx) error {
	idParam := c.Params("id")
	postID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "ID artikel tidak valid"})
	}

	var input struct {
		UserID uint `json:"user_id"`
	}
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Input user_id dibutuhkan"})
	}

	if err := h.Service.DeletePost(uint(postID), input.UserID); err != nil {
		if err.Error() == "akses ditolak: kamu hanya bisa menghapus artikelmu sendiri" {
			return c.Status(403).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "Artikel berhasil dihapus"})
}

func (h *PostHandler) GetAll(c *fiber.Ctx) error {
	posts, err := h.Repo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	return c.JSON(fiber.Map{"data": posts})
}

func (h *PostHandler) GetBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	post, err := h.Repo.FindBySlug(slug)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Artikel tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"data": post})
}