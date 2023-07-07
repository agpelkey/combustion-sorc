package data

import (
	"database/sql"
)

type Item struct {
	ID    int `json:"id"`
	Name  int `json:"name"`
	Price int `json:"price"`
}

type Models struct {
	Items PostgresDBRepo
}

func NewModels(db *sql.DB) Models {
	return Models{
		Items: PostgresDBRepo{DB: db},
	}
}
