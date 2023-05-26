package model

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	Id        int       `json:"id"`
	Author    string    `json:"author"`
	UserId    int       `json:"user_id"`
	Title     string    `json:"caption"`
	Content   string    `json:"content"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt
}
