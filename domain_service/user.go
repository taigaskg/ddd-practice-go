package domainservice

import (
	"database/sql"
	"ddd-practice-go/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserService struct{}

func (us *UserService) Exists(u *model.User) bool {
	db, err := sql.Open(
		driverName,
		fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, databaseName),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	username := u.GetName()
	rows, err := db.Query("SELECT * FROM users WHERE name = ?", username.ToString())
	if err != nil {
		log.Fatal(err)
	}

	var cnt int
	for rows.Next() {
		cnt++
	}
	defer rows.Close()
	return cnt > 0
}

const (
	driverName   = "mysql"
	user         = "root"
	password     = "password"
	host         = "db"
	port         = 3306
	databaseName = "api-covid19"
)
