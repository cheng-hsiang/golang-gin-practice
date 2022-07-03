package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID         uuid.UUID `json:"id" gorm:"type:char(36);primary_key"`
	UserId     uint      `json:"user_id" gorm:"not null"`
	CategoryId uint      `json:"category_id" gorm:"not null"`
	Title      string    `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg    string    `json:"head_img"`
	Content    string    `json:"content" gorm:"type:text;not null"`
	Category   *Category
}

func (post *Post) BeforeCreate(tx *gorm.DB) error {
	post.ID = uuid.New()
	return nil
}
