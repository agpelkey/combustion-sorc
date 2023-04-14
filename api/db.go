package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const dbtimeout = time.Second * 5

// make storage interface
type InventoryStorage interface {
	AddItem(item *InventoryItem) error
}

// make Postgres struct
type PostgresDB struct {
	db *sql.DB
}

// make Postgres connection function
func NewPostgresDB() (*PostgresDB, error) {

	//var passwd = os.Getenv("$POSTGRES_PASSWORD")

	connstr := "user=postgres dbname=postgres password=inventory sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return &PostgresDB{
		db: db,
	}, nil
}

// make Init function to create tables
func (s *PostgresDB) Init() error {
	return s.createInventoryTable()
}

// make function to generate tables
func (s *PostgresDB) createInventoryTable() error {
	stmt := `CREATE TABLE IF NOT EXISTS inventory (
			id SERIAL PRIMARY KEY,
			name VARCHAR(50),
			amount INTEGER
		);`

	_, err := s.db.Exec(stmt)
	return err
}

// function to add items to database
func (s *PostgresDB) AddItem(item *InventoryItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbtimeout)
	defer cancel()

	query := `INSERT INTO inventory (name, amount) VALUES ($1, $2)`

	_, err := s.db.QueryContext(ctx, query,
		item.Name,
		item.Amount,
	)
	if err != nil {
		return err
	}

	return nil

}
