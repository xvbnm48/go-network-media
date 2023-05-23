package user

import (
	"github.com/xvbnm48/go-network-media/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindUserById(id int) (model.User, error)
	Save(user model.User) (model.User, error)
	DeleteUser(id int) error
	FindByEmail(email string) (model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUserById(id int) (model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Save(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) DeleteUser(id int) error {
	err := r.db.Delete(&model.User{}, id).Error
	if err != nil {
		return err
	}

	return nil

}

func (r *repository) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
