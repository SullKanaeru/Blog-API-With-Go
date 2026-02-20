package model

import "time"

type PostStatus string

const (
	StatusDraft     PostStatus = "draft"
	StatusPublished PostStatus = "published"
)

type Post struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Slug      string     `json:"slug"`
	Content   string     `json:"content"`
	Status    PostStatus `json:"status"`
	UserID    uint       `json:"user_id"`
	
	User      User       `json:"author" gorm:"foreignKey:UserID"` 
	
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}