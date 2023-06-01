package reaction

import (
	"github.com/xvbnm48/go-network-media/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	AddLike(userId int, postId int) (int, error)
	AddUnlike(userId int, postId int) (int, error)
	RemoveLike(userId int, postId int) (int, error)
	RemoveUnlike(userId int, postId int) (int, error)
	CountLike(postId int) (int64, error)
	CountUnlike(postId int) (int64, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddLike(userId int, postId int) (int, error) {
	var reaction = model.Reaction{}
	err := r.db.Debug().Model("reaction").Create(&reaction).Error
	if err != nil {

		return 0, err
	}

	return 1, nil
}
