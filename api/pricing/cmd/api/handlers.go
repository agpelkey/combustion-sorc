package main

import (
	"net/http"
)

func (app *application) handleGetPrice(w http.ResponseWriter, r *http.Request) {
}

/*
func (s *APIServer) handleAddItem(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleGetPrice(w http.ResponseWriter, r *http.Request) error {

	params := mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Fatal(err)
	}

	var item Item

	coll := s.db.Client().Database("ItemPrices").Collection("price")

	err = coll.FindOne(context.TODO(), Item{ID: id}).Decode(&item)
	if err != nil {
		log.Println(err)
	}

	json.NewEncoder(w).Encode(item)

	return
}

func (s *APIServer) handleAddItem(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {

		// initilialize Item struct
		item := Item{}

		// create database connection
		coll := s.db.Client().Database("ItemPrices").Collection("prices")

		// create variable to read document implementation
		payload := json.NewDecoder(r.Body).Decode(&item)

		// insert the payload into the DB
		result, err := coll.InsertOne(context.TODO(), payload)
		if err != nil {
			panic(err)
		}

		log.Println("Inserted document into database with _id: \n", result.InsertedID)

		json.NewEncoder(w).Encode(result)
	}

	return nil
}
*/

// for future implementation. Will need to figure out how to gate this behind admin access so not anyone can just alter the price.
// func (s *APIServer) handleUpdateItemPrice(w http.ResponseWriter, r *http.Request) error {
//	 return nil
// }