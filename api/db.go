package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// make storage interface
type InventoryStorage interface {
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
