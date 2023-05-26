package user

import "github.com/xvbnm48/go-network-media/model"

type userFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserFormatterWithFriends struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	//Followers int64  `json:"followers"`
}

func FormatUserWithFriends(user model.User) UserFormatterWithFriends {
	formatter := UserFormatterWithFriends{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		//Followers: user.Friends,
	}

	return formatter
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

//
//func FormatUserWithFriends(user model.User) {
//	formatter := userFormatterWithFriends{
//		ID:        user.Id,
//		Name:      user.Name,
//		Email:     user.Email,
//		Followers: user.Friends,
//	}
//}
