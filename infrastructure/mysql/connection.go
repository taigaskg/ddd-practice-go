package mysql

import (
	"database/sql"
	"fmt"
)

const (
	driverName   = "mysql"
	user         = "root"
	password     = "password"
	host         = "db"
	port         = 3306
	databaseName = "api-covid19"
)

func OpenDB() (*sql.DB, error) {
	db, err := sql.Open(
		driverName,
		fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=true", user, password, host, port, databaseName),
	)
	if err != nil {
		return nil, err
	}
	return db, nil
}
