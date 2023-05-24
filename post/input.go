package post

type PostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type UpdatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
