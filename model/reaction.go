package model

type Reaction struct {
	UserId int `json:"user_id"`
	User   User
	PostId int `json:"post_id"`
	Post   Post
	Status bool
}
