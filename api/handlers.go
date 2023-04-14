package main

import (
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
	return nil
}

func (s *APIServer) handleUpdateInventory(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteInventory(w http.ResponseWriter, r *http.Request) error {
	return nil
}
