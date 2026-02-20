package repository

import (
	"blog_api/internal/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) Create(post *model.Post) error {
	return r.DB.Create(post).Error
}

func (r *PostRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.DB.Preload("User").Find(&posts).Error
	return posts, err
}

func (r *PostRepository) FindByID(id uint) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("User").Where("id = ?", id).First(&post).Error
	return &post, err
}

func (r *PostRepository) FindBySlug(slug string) (*model.Post, error) {
	var post model.Post
	err := r.DB.Preload("User").Where("slug = ?", slug).First(&post).Error
	return &post, err
}

func (r *PostRepository) Update(post *model.Post) error {
	return r.DB.Save(post).Error
}

func (r *PostRepository) Delete(id *model.Post) error {
	return r.DB.Delete(&model.Post{}, id).Error
}