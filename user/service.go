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
