package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	case r.Method == "PATCH":
		s.handleUpdateInventory(w, r)
	case r.Method == "DELETE":
		s.handleDeleteInventory(w, r)
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

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return err
	}

	if err := s.db.UpdateItem(id); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (s *APIServer) handleDeleteInventory(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// helper function to write json
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
