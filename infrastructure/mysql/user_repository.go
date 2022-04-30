package mysql

import (
	"database/sql"
	"ddd-practice-go/domain/model"

	_ "github.com/go-sql-driver/mysql"
)

type userRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) FetchByName(name string) ([]*model.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var id string
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return nil, err
		}

		userId, err := model.NewUserID(id)
		if err != nil {
			return nil, err
		}

		userName, err := model.NewUserName(name)
		if err != nil {
			return nil, err
		}

		user, err := model.NewUser(*userId, *userName)
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) Register(user *model.User) error {
	stmt, err := r.DB.Prepare("INSERT INTO users(id, name) VALUES (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(user.ID(), user.Name())
	if err != nil {
		return err
	}
	return nil
}
