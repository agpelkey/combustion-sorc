package main

import (
	"context"
	"flag"
	"os"

	"github.com/agpelkey/combustion-sorc/internal/data"
	"github.com/agpelkey/combustion-sorc/internal/jsonlog"
	"github.com/jackc/pgx/v5/pgxpool"
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

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("PRICING_DB_DSN"))
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer dbpool.Close()

	logger.PrintInfo("database connection established", nil)
	app := &application{
		config: cfg,
		logger: logger,
		items:  data.NewItems(dbpool),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
