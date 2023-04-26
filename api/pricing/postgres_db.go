package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// Create interface for DB connections
type PriceStorage interface {
}

// Create Postgres DB struct
type PostgresDB struct {
	DB *sql.DB
}

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
