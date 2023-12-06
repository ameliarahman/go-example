package database

import (
	"database/sql"
	"goexample/internal/config"

	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {

	user := config.UserPostgres
	password := config.PasswordPostgres
	host := config.HostPostgres
	dbname := config.NamePostgres
	sslmode := "disable"

	connectionConfig := "user=" + user + " password=" + password + " host=" + host + " dbname=" + dbname + " sslmode=" + sslmode
	db, err := sql.Open("postgres", connectionConfig)
	if err != nil {
		return nil, err
	}

	return db, nil
}
