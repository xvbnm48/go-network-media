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

type AllPost struct {
	Id        int       `gorm:"column:id;primaryKey" json:"id"`
	Author    string    `gorm:"column:author" json:"author"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	Title     string    `gorm:"column:title" json:"caption"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

//func (AllPost) Posts() string {
//	return "posts"
//}
