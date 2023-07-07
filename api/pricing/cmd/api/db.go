package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil

}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.config.db.dsn)
	if err != nil {
		return nil, err
	}

	app.logger.PrintInfo("Connected to Postgres instance", nil)
	return connection, nil
}
