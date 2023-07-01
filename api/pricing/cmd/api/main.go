package main

import (
	"database/sql"
	"flag"
	"os"

	"github.com/agpelkey/combustion-sorc/internal/data"
	"github.com/agpelkey/combustion-sorc/internal/jsonlog"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
		// Have not yet found a way to incorporate this with pgxpool
		//maxOpenConns int
		//maxIdleConns int
		//maxIdleTime  string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	items  data.Models
}

func main() {

	var cfg config

	// Read in the value for port and env
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("PRICING_DB_DSN"), "PostgreSQL DB DSN")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	/*
		dbConn, err := NewPostgresDB()
		if err != nil {
			log.Fatal(err)
		}

		if err := dbConn.createPricingTable(); err != nil {
			log.Fatal(err)
		}
	*/
	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database connection *soon* to be established", nil)
	app := &application{
		config: cfg,
		logger: logger,
		items:  data.NewItems(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	// create db connection pool
	db, err := sql.Open("pgxpool", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil

}
