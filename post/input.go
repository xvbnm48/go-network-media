package post

import "github.com/xvbnm48/go-network-media/model"

type PostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type UpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
	User    model.User
}

type GetPostDetailInput struct {
	Id int `uri:"id" binding:"required"`
}
