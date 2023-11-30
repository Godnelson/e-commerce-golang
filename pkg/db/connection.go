package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type Connection struct{}

func (d *Connection) New() *sqlx.DB {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=root password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
