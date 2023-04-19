package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5"
)

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Server Running")
	return nil
}

func (s *APIServer) handleInventory(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == "GET":
		s.handleGetInventory(w, r)
	case r.Method == "POST":
		s.handleAddInventory(w, r)
	}

	return nil
}

func (s *APIServer) handleGetInventory(w http.ResponseWriter, r *http.Request) error {
	obj, err := s.db.GetAllItems()
	if err != nil {
		log.Fatal(err)
	}

	return writeJSON(w, http.StatusOK, obj)
}

func (s *APIServer) handleAddInventory(w http.ResponseWriter, r *http.Request) error {
	// create item variable
	var item InventoryItem

	// decode request body into item variable
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		return err
	}

	// call APIServer db function to add item to PostgresDB
	newItem := s.db.AddItem(&item)
	if err != nil {
		return err
	}

	// return success
	return writeJSON(w, http.StatusOK, newItem)

}

func (s *APIServer) handleUpdateInventory(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "PUT" {
		fmt.Errorf("Method %s not allowed")
	}

	id := chi.URLParam(r, "id")

	requestID, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.db.UpdateItem(requestID); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *APIServer) handleDeleteInventory(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "DELETE" {
		id := chi.URLParam(r, "id")

		strID, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		if err := s.db.DeleteItem(strID); err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// helper function to write json
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
