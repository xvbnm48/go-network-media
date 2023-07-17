package vote

import (
	"github.com/xvbnm48/go-network-media/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	CountVote(userID int, Id int) bool
	AddVote(userID int, id int)
	ReadVote(id string) []string
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//	func (r *repository) AddLike(userId int, postId int) (int, error) {
//		var reaction = model.Votes{}
//		err := r.db.Debug().Model("vote").Create(&reaction).Error
//		if err != nil {
//
//			return 0, err
//		}
//
//		return 1, nil
//	}
func (r *repository) CountVote(userID int, Id int) bool {
	var count int
	r.db.Debug().Raw("SELECT COUNT(*) FROM votes where user_id = $1 and id = $2", userID, Id).Scan(&count)

	switch count {
	case 0:
		return false
	default:
		return true
	}
}
func (r *repository) AddVote(userID int, id int) {
	var query string
	voted := r.CountVote(userID, id)
	switch voted {
	case false:
		query = "INSERT INTO votes (user_id, id) VALUES ($1, $2)"
	case true:
		query = "DELETE FROM votes WHERE user_id = $1 and id = $2"
	}

	if err := r.db.Exec(query, userID, id); err != nil {
		panic(err)
	}
	// this is for the first time vote
}

func (r *repository) ReadVote(id string) []string {
	//var voter []string
	//rows, err := r.db.Exec("SELECT name FROM users WHERE id IN (SELECT user_id FROM votes WHERE id = $1)", id).Rows()
	//if err != nil {
	//	log.Println(err)
	//	return nil
	//}
	//
	//defer rows.Close()
	//for rows.Next() {
	//	var name string
	//	if err := rows.Scan(&name); err != nil {
	//		log.Println(err)
	//		return nil
	//	}
	//	voter = append(voter, name)
	//}
	//return voter
	var voters []model.User
	err := r.db.Table("users").Select("name").Joins("JOIN votes ON users.id = votes.user_id").Where("votes.id = ?", id).Find(&voters).Error
	if err != nil {
		return nil
	}

	var votersName []string
	for _, voter := range voters {
		votersName = append(votersName, voter.Name)
	}

	return votersName
}
