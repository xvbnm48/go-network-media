package user

import "github.com/xvbnm48/go-network-media/model"

type userLoginFormatter struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type PostGet struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}

type UserFormatterWithFriends struct {
	ID        int             `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Followers int64           `json:"followers"`
	Following int64           `json:"following"`
	Post      []model.AllPost `json:"posts"`
}

func FormatUserWithFriends(user model.User) UserFormatterWithFriends {
	formatter := UserFormatterWithFriends{
		ID:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Followers: user.Followers,
		Following: user.Following,
		Post:      user.Posts,
	}

	return formatter
}

func FormatUser(user model.User, token string) userLoginFormatter {
	formatter := userLoginFormatter{
		ID:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return formatter
}
