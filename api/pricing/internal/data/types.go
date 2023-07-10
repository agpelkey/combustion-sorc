package data

import (
	"database/sql"
)

type Item struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
}

type Models struct {
	Items PostgresDBRepo
}

func NewModels(db *sql.DB) Models {
	return Models{
		Items: PostgresDBRepo{DB: db},
	}
}
