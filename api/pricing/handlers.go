package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (s *APIServer) handleGetPrice(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleAddItem(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		fmt.Errorf("Error method %s not allowed", r.Method)
	}

	// initilialize Item struct
	item := Item{}

	// create variable to read document implementation
	payload := json.NewDecoder(r.Body).Decode(&item)

	// create database connection
	coll := s.db.Client().Database("ItemPrices").Collection("prices")

	// insert the payload into the DB
	result, err := coll.InsertOne(context.TODO(), payload)
	if err != nil {
		panic(err)
	}

	log.Println("Inserted document into database with _id: %v\n", result.InsertedID)

	return nil
}

// for future implementation. Will need to figure out how to gate this behind admin access so not anyone can just alter the price.
// func (s *APIServer) handleUpdateItemPrice(w http.ResponseWriter, r *http.Request) error {
//	 return nil
// }
