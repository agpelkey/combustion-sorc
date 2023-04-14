package main

import "log"

func main() {

	// create db connection
	db, err := NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.createInventoryTable(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8080", db)

	server.Run()
}
