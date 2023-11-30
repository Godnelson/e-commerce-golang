package category

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Query struct {
	db *sqlx.DB
}

func (q *Query) GetAll() []Category {
	var categories []Category
	err := q.db.Select(&categories, "SELECT * FROM category")
	if err != nil {
		log.Fatal(err)
	}
	return categories
}
