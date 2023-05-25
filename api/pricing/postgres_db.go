package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Create interface for DB connections
type PriceStorage interface {
    GetItemByID(id int) ([]*Item, error)
}

// Create Postgres DB struct
type PostgresDB struct {
	DB *sql.DB
}

const dbtimeout = 10 * time.Second

// Make postgres connection to DB
func NewPostgresDB() (*PostgresDB, error) {

	connStr := "user=postgres dbname=postgres password=pricing sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresDB{
		DB: db,
	}, nil
}

// Function to Init DB
func (p *PostgresDB) Init() error {
	return p.createPricingTable()
}

// Create DB table
func (p *PostgresDB) createPricingTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS pricing (
			id SERIAL PRIMARY KEY,
			price INTEGER
		);`

	_, err := p.DB.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (p *PostgresDB) GetItemByID(id int) ([]*Item, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbtimeout) 
    defer cancel()

    query := `SELECT id, price WHERE id = $1`

    rows, err := p.DB.QueryContext(ctx, query, id)
    if err != nil {
        log.Fatal(err)
    }

    var items []*Item

    for rows.Next() {
        var item Item
        err = rows.Scan(
            &item.ID,
            &item.Price,
        )
        if err != nil {
            return nil, err 
        }

        items = append(items, &item)
    }

    return items, nil
}














