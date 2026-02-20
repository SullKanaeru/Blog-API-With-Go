package service

import (
	"errors"
	"strings"
	"blog_api/internal/model"
	"blog_api/internal/repository"
)

type PostService struct {
	Repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(input *model.Post) error {
	if input.Title == "" || input.Content == "" {
		return errors.New("judul dan konten tidak boleh kosong")
	}

	if input.Slug == "" {
		input.Slug = strings.ToLower(strings.ReplaceAll(input.Title, " ", "-"))
	}

	if input.Status == "" {
		input.Status = model.StatusDraft
	}

	return s.Repo.Create(input)
}

func (s *PostService) UpdatePost(postID uint, userID uint, req model.Post) (*model.Post, error) {
	existingPost, err := s.Repo.FindByID(postID)
	if err != nil {
		return nil, errors.New("artikel tidak ditemukan")
	}

	if existingPost.UserID != userID {
		return nil, errors.New("akses ditolak: kamu hanya bisa mengedit artikelmu sendiri")
	}

	if req.Title != "" {
		existingPost.Title = req.Title
		existingPost.Slug = strings.ToLower(strings.ReplaceAll(req.Title, " ", "-"))
	}
	if req.Content != "" {
		existingPost.Content = req.Content
	}
	if req.Status != "" {
		existingPost.Status = req.Status
	}

	if err := s.Repo.Update(existingPost); err != nil {
		return nil, errors.New("gagal menyimpan pembaruan artikel")
	}

	return existingPost, nil
}

func (s *PostService) DeletePost(postID uint, userID uint) error {
	existingPost, err := s.Repo.FindByID(postID)
	if err != nil {
		return errors.New("artikel tidak ditemukan")
	}

	if existingPost.UserID != userID {
		return errors.New("akses ditolak: kamu hanya bisa menghapus artikelmu sendiri")
	}

	return s.Repo.Delete(existingPost)
}