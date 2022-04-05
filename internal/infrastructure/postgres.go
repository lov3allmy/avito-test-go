package infrastructure

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	host     = "localhost"
	port     = 5436
	username = "postgres"
	password = "qwerty"
	dbname   = "avito_test_go"
)

func ConnectToPostgres() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname))
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
