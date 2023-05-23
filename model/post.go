package model

type Post struct {
	Id      string `json:"id"`
	Caption string `json:"caption"`
	UserID  string `json:"user_id"`
}
