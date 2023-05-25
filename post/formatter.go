package post

import "github.com/xvbnm48/go-network-media/model"

type PostFormatter struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}

func FormatPost(post model.Post) PostFormatter {
	postFormat := PostFormatter{}
	postFormat.Title = post.Title
	postFormat.Content = post.Content
	postFormat.Author = post.Author
	postFormat.CreatedAt = post.CreatedAt.Format("2006-01-02 15:04:05")

	return postFormat
}

func FormatPosts(posts []model.Post) []PostFormatter {
	if len(posts) == 0 {
		return []PostFormatter{}
	}
	var postsFormatter []PostFormatter

	for _, post := range posts {
		postFormatter := FormatPost(post)
		postsFormatter = append(postsFormatter, postFormatter)
	}

	return postsFormatter
}
