package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type APIServer struct {
	listenAddr string
	db         *mongo.Database
}

func NewAPIServer(listenAddr string, mg *MongoInstance) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         mg.DB,
	}
}

// adapter to allow ordinary functions to be used as HTTP Handlers
// will only really be used in the makeHTTPHandler function below
type apifunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *APIServer) Run() {

	mux := http.NewServeMux()

	mux.Handle("/api/warehouse/pricing", makeHTTPHandler(s.handlePricing))

	log.Println("Starting API server on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, mux)
}
