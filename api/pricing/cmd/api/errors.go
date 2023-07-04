package main

import (
	"log"
	"net/http"
)

// Some basic error handlers. This file can be built out with many more functions
// depending on the size and scope of the application

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {

	env := envelope{"error": message}

	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		log.Println(r, err)
		w.WriteHeader(500)
	}
}

// used when the application encounters error at run time.
// Method will log the error message and then user errorResponse()
// to send a 500 Internal Server Error status code and JSON repsonse
// to the client.
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(r, err)

	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// function to return a 404 Not Found Response status code and JSON response to the clietn
func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}
