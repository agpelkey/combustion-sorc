package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	listenAddr string
	db         InventoryStorage
}

func NewAPIServer(listenAddr string, db InventoryStorage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         db,
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *APIServer) Run() {

	router := chi.NewRouter()

	router.HandleFunc("/api/warehouse/inventory", makeHTTPHandler(s.handleInventory))
	router.HandleFunc("/api/warehouse/inventory/delete/{id}", makeHTTPHandler(s.handleDeleteInventory))
	router.HandleFunc("/api/warehouse/inventory/update/{id}", makeHTTPHandler(s.handleUpdateInventory))

	log.Println("Starting server on port ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
