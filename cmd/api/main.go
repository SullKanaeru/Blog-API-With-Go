package main

import (
	"blog_api/internal/config"
	"blog_api/internal/handler"
	"blog_api/internal/middleware"
	"blog_api/internal/repository"
	"blog_api/internal/service"
	database "blog_api/pkg/database/migrations"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	db := database.ConnectDB()

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService, postRepo)

	app := fiber.New()
	app.Use(logger.New())

	api := app.Group("/api/posts", middleware.RequireAuth())
	api.Post("/", middleware.AllowRoles("author"), postHandler.CreatePost)
	api.Put("/:id", postHandler.UpdatePost)
	api.Delete("/:id", postHandler.DeletePost)
	api.Get("/", postHandler.GetAll)
	api.Get("/:slug", postHandler.GetBySlug)

	userRepo := repository.NewUserRepository(db)

	authService := service.NewAuthService(userRepo, config.GetEnv("JWT_SECRET", "default_secret"))
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(authService, userRepo)

	userRoutes := app.Group("/api/users")
	userRoutes.Post("/register", authHandler.Register)
	userRoutes.Post("/login", authHandler.Login)
	userRoutes.Post("/logout", authHandler.Logout)
	userRoutes.Get("/", userHandler.GetAllUsers)

	protectedUsers := app.Group("/api/users", middleware.RequireAuth())
	protectedUsers.Get("/me", userHandler.GetProfile)

	port := config.GetEnv("PORT", "3000")
	app.Listen(":" + port)
}
