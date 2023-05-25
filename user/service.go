package user

import (
	"errors"
	"github.com/xvbnm48/go-network-media/model"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (model.User, error)
	LoginUser(input LoginUserInput) (model.User, error)
	GetUserById(ID int) (model.User, error)
	IsEmailAvailable(email string) (bool, error)
	FollowFriends(userId int, friendId int) (int, error)
	//UnfollowFriends(userId int, friendId int) (int, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) RegisterUser(input RegisterUserInput) (model.User, error) {
	user := model.User{}
	user.Name = input.Name
	user.Email = input.Email
	s.IsEmailAvailable(input.Email)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	NewUser, err := s.repo.Save(user)
	if err != nil {
		return user, err
	}

	return NewUser, nil
}

func (s *service) LoginUser(input LoginUserInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("No user found on that email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) GetUserById(ID int) (model.User, error) {
	user, err := s.repo.FindUserById(ID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("No user found on that ID")
	}

	return user, nil
}

func (s *service) IsEmailAvailable(email string) (bool, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.Id == 0 {
		return true, errors.New("No user found on that email")
	}

	return false, nil
}

func (s *service) FollowFriends(userId int, friendId int) (int, error) {
	_, err := s.repo.FindUserById(userId)
	if err != nil {
		return 0, err
	}
	_, err = s.repo.FindUserById(friendId)
	if err != nil {
		return 0, err
	}

	if userId == friendId {
		return 0, errors.New("You can't follow yourself")
	}

	_, err = s.repo.Follow(userId, friendId)
	if err != nil {
		return 0, err
	}

	return friendId, nil
}
