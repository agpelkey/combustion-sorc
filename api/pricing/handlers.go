package main

import "net/http"

func (s *APIServer) handlePricing(w http.ResponseWriter, r *http.Request) error {
	switch {
	case r.Method == "GET":
		// s.handleGetPrice(w, r)
		return nil
	case r.Method == "PUT":
		//s.handleUpdateItemPrice(w, r)
		return nil

	}

	return nil
}

func (s *APIServer) handleGetPrice(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateItemPrice(w http.ResponseWriter, r *http.Request) error {
	return nil
}
