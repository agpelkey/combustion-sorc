package main

import "log"

func main() {

	dbConn, err := NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := dbConn.createPricingTable(); err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":8081", dbConn)

	server.Run()
}
