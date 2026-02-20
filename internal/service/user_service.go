package service

import (
	"blog_api/internal/model"
	"blog_api/internal/repository"
	"errors"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(input model.User) error {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("nama, email, dan password wajib diisi")
	}

	role := input.Role
	if role == "" {
		role = "viewer"
	}

	if input.Role != "author" && input.Role != "" {
		return errors.New("role tidak valid, gunakan: author atau viewer")
	}

	return s.Repo.CreateUser(&input)
}
