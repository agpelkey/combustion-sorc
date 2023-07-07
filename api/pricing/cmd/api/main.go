package main

import (
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
	models data.Models
}

func main() {

	var cfg config

	// Read in the value for port and env
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("PRICING_DB_DSN"), "PostgreSQL DB DSN")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg.db.dsn)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	defer db.Close()

	logger.PrintInfo("database connection established", nil)
	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
