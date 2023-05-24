package post

import "github.com/xvbnm48/go-network-media/model"

type service struct {
	repository Repository
}

type Service interface {
	GetAllPost() ([]model.Post, error)
	CreatePost(input PostInput) (model.Post, error)
	DeletePost(id int) error
	UpdatePost(input UpdatePost) (model.Post, error)
}

func NewServicePost(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreatePost(input PostInput) (model.Post, error) {
	inputPost := model.Post{}
	inputPost.Title = input.Title
	inputPost.Content = input.Content
	inputPost.Author = input.Author

	newPost, err := s.repository.CreatePost(inputPost)
	if err != nil {
		return newPost, err
	}

	return newPost, nil
}

func (s *service) UpdatePost(input UpdatePost) (model.Post, error) {
	inputPost := model.Post{}
	inputPost.Title = input.Title
	inputPost.Content = input.Content

	updatePost, err := s.repository.UpdatePost(inputPost)
	if err != nil {
		return updatePost, err
	}

	return updatePost, nil
}

func (s *service) DeletePost(id int) error {
	post := s.repository.DestroyPost(id)
	if post != nil {
		return post
	}

	return nil
}

func (s *service) GetAllPost() ([]model.Post, error) {
	posts, err := s.repository.FindAll()
	if err != nil {
		return posts, err
	}

	return posts, nil
}
