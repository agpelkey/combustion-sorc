package main

import (
	"flag"
	"os"

	"github.com/agpelkey/combustion-sorc/internal/jsonlog"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn          string
		maxOpenConns int
		maxIdleConns int
		maxIdleTime  string
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
}

func main() {

	var cfg config

	// Read in the value for port and env
	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

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

	logger.PrintInfo("database connection *soon* to be established", nil)
	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}
