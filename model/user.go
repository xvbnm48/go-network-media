package model

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Followers int64     `json:"followers"`
	Following int64     `json:"following"`
	Posts     []AllPost `json:"posts"`
	Friends   []User    `gorm:"many2many:friendships;association_jointable_foreignkey:friend_id"`
}

type UserWithFriends struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Followers int64  `json:"followers"`
}
