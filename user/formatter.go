package user

import "github.com/xvbnm48/go-network-media/model"

type userFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func FormatUser(user model.User, token string) userFormatter {
	formatter := userFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}
