package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Item struct {
	ID    int `json:"id"`
	Name  int `json:"name"`
	Price int `json:"price"`
}

// Create interface for DB connections
type PriceStorage interface {
	GetItemByID(id int) (*Item, error)
}

// Create Postgres DB struct
type ItemModel struct {
	DB *pgxpool.Pool
}

func (i ItemModel) GetItemByID(id int64) (*Item, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, name, price
						FROM item_prices
						WHERE id = $1`

	var item Item

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := i.DB.QueryRow(ctx, query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Price,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &item, nil
}
