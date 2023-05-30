package post

import (
	"github.com/xvbnm48/go-network-media/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	FindAll() ([]model.Post, error)
	CreatePost(post model.Post) (model.Post, error)
	UpdatePost(post model.Post) (model.Post, error)
	DestroyPost(id int) error
	FindPostById(id int) (model.Post, error)
	FindAllPosts(userId int, post []model.AllPost) ([]model.AllPost, error)
}

func NewPostRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePost(post model.Post) (model.Post, error) {
	err := r.db.Create(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) UpdatePost(post model.Post) (model.Post, error) {
	err := r.db.Save(&post).Error
	if err != nil {
		return post, err
	}

	return post, nil
}

func (r *repository) DestroyPost(id int) error {
	err := r.db.Delete(&model.Post{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return posts, err
	}

	return posts, nil

}

func (r *repository) FindPostById(id int) (model.Post, error) {
	//post := &model.Post{}
	//err := r.db.Preload("User").First(&post, id).Error
	//if err != nil {
	//	return post, err
	//}
	//
	//return post, nil
	post := model.Post{}
	err := r.db.Preload("User").Where("id = ?", id).First(&post).Error
	if err != nil {
		return post, err
	}
	return post, nil
}

func (r *repository) FindAllPosts(userId int, post []model.AllPost) ([]model.AllPost, error) {
	var posts = post
	err := r.db.Debug().Raw("SELECT * FROM posts WHERE user_id = ? AND posts.deleted_at is NULL", userId).Scan(&posts).Error
	if err != nil {
		return posts, err
	}
	return posts, nil
}
