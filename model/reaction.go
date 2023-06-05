package model

// models vote
type Votes struct {
	id     int `json:"post_id"`
	Post   Post
	UserId int `json:"user_id"`
	User   User
}
