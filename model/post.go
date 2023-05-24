package model

import (
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"caption"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
