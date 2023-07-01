package data

import (
	"database/sql"
)

type Models struct {
	Items ItemModel
}

func NewItems(db *sql.DB) Models {
	return Models{
		Items: ItemModel{DB: db},
	}
}
