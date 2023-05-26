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
	Follow(userId int, friendId int) (int, error)
	Unfollow(userId int, friendId int) (int, error)
	CountFollowers(id int) (int64, error)
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

func (r *repository) Follow(userId int, friendId int) (int, error) {
	user := model.User{}
	friend, err := r.FindUserById(friendId)
	if err != nil {
		return userId, err
	}

	err = r.db.Preload("Friends").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return userId, err
	}

	err = r.db.Model(&user).Association("Friends").Append(&friend)
	if err != nil {
		return userId, err
	}

	return friendId, nil
}

func (r *repository) Unfollow(userId int, friendId int) (int, error) {
	user := model.User{}
	friend, err := r.FindUserById(friendId)
	if err != nil {
		return userId, err
	}
	err = r.db.Preload("Friends").Where("id = ?", userId).First(&user).Error
	if err != nil {
		return userId, err
	}

	err = r.db.Model(&user).Association("Friends").Delete(&friend)

	return friendId, nil
}
func (r *repository) CountFollowers(id int) (int64, error) {
	var count int64
	err := r.db.Table("friendships").Where("friend_id = ?", id).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
