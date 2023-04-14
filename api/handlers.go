package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) handleHome(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Server Running")
	return nil
}

func (s *APIServer) handleInventory(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == "GET":
		s.handleGetInventory(w, r)
		return nil
	case r.Method == "POST":
		s.handleAddInventory(w, r)
	case r.Method == "PUT":
		s.handleUpdateInventory(w, r)
	case r.Method == "DELETE":
		s.handleDeleteInventory(w, r)
	}

	return fmt.Errorf("Error, method %s not allowed", r.Method)
}

func (s *APIServer) handleGetInventory(w http.ResponseWriter, r *http.Request) error {
	return nil
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
