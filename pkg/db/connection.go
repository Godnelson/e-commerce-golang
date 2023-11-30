package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Connection struct {
	Client *sqlx.DB
}

func New() Connection {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=root password=%s dbname=%s sslmode=disable", os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME")))
	if err != nil {
		log.Fatal(err)
	}
	return Connection{db}
}
