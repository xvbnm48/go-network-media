package post

import (
	"errors"
	"fmt"
	"github.com/xvbnm48/go-network-media/model"
)

type service struct {
	repository Repository
}

type Service interface {
	GetAllPost() ([]model.Post, error)
	CreatePost(input PostInput) (model.Post, error)
	DeletePost(id int) error
	//UpdatePost(InputID GetPostDetailInput, input UpdatePost) (model.Post, error)
	//UpdatePost(post PostInput, postId int, userId int) (model.Post, error)
	UpdatePost(inputID GetPostDetailInput, inputPost UpdatePost) (model.Post, error)
}

func NewServicePost(repository Repository) *service {
	return &service{repository: repository}
}

//inputPost.Id = input.User.Id

func (s *service) CreatePost(input PostInput) (model.Post, error) {
	inputPost := model.Post{}
	inputPost.Title = input.Title
	inputPost.Content = input.Content
	inputPost.Author = input.Author
	inputPost.User.Id = input.User.Id
	newPost, err := s.repository.CreatePost(inputPost)
	if err != nil {
		return newPost, err
	}

	return newPost, nil
}

func (s *service) UpdatePost(inputID GetPostDetailInput, inputPost UpdatePost) (model.Post, error) {
	post, err := s.repository.FindById(inputID.Id)
	fmt.Println("isi post dari hasil data findbyid: ", post)
	if err != nil {
		return post, err
	}
	fmt.Println("isi post user id: ", post.User, "dan isi dari input post user id: ", inputPost.User.Id)
	if post.User.Id != inputPost.User.Id {
		return post, errors.New("you are not the owner of this post")
	}
	post.Title = inputPost.Title
	post.Content = inputPost.Content
	post.Author = inputPost.Author

	updatePost, err := s.repository.UpdatePost(post)
	if err != nil {
		return updatePost, err
	}

	return updatePost, nil
}

//func (s *service) UpdatePost(InputID GetPostDetailInput, input UpdatePost) (model.Post, error) {
//	oldPost := model.Post{}
//	updatedPost, err := s.repository.FindById(oldPost)
//	if err != nil {
//		return updatedPost, err
//	}
//
//	//	post, err := s.repository.FindById(InputID.Id)
//	//	if err != nil {
//	//		return post, err
//	//	}
//	//
//	//	if post.Id != input.id {
//	//		return post, errors.New("post not found")
//	//	}
//	//	inputPost := model.Post{}
//	//	inputPost.Title = input.Title
//	//	inputPost.Content = input.Content
//	//	inputPost.Author = input.Author
//	//
//	//	updatePost, err := s.repository.UpdatePost(inputPost)
//	//	if err != nil {
//	//		return updatePost, err
//	//	}
//	//
//	//	return updatePost, nil
//}

//func (s *service) UpdatePost(post PostInput, postId int, userId int) (model.Post, error) {
//	// postId from uri, userID from token
//	oldPost, err := s.repository.FindById(postId)
//	if err != nil {
//		return oldPost, err
//	}
//	fmt.Println("oldp post user ", oldPost.User, "user id dari context", userId, "oldpost id", oldPost.Id)
//	if oldPost.User.Id != userId {
//		return oldPost, errors.New("you are not the owner of this post")
//	}
//
//	oldPost.Title = post.Title
//	oldPost.Content = post.Content
//	oldPost.Author = post.Author
//	updatedPost, err := s.repository.UpdatePost(oldPost)
//	if err != nil {
//		return updatedPost, err
//	}
//
//	return updatedPost, nil
//}

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
