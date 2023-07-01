package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	// Insert routes here
	router.HandlerFunc(http.MethodGet, "/api/pricing", app.handleGetPrice)

	return router
}
