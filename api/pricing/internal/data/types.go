package data

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Items ItemModel
}

func NewItems(db *pgxpool.Pool) Models {
	return Models{
		Items: ItemModel{DB: db},
	}
}
