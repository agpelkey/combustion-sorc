package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// enevelope will be used throughout this application to "write" messages
type envelope map[string]interface{}

// function to read JSON from the client. This is a good place to triage any errors
// from the incoming request. Error handling can be built out depending on size
// and scope of the application
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	// use maxBytesReader to limit the size of the request body to 1 MB
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// Calling DisallowUnknownFields() means that an error
	// will be returned if the JSON from the client includes any field which cannot
	// be mapped to the target destination
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	// decode request body into target destination
	err := dec.Decode(dst)
	if err != nil {
		// this is where we can do some meaty error triaging.
		// We could use a switch statement to look for error such as
		// syntaxErrors, UnmarshalTypeErrors, and InvalidUnmarshalErrors.
		// For this service the triage will not need to be as built out
		return err
	}

	// calling decode one more time using a pointer to an anonymous struct.
	// This will return an io.EOF if the request body only contains a single
	// JSON value. Otherwise, we return our custom error
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value")
	}

	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, header http.Header) error {

	// Encode the data to JSON
	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// we can append a new line to make it easier to view in terminal.
	// This is strictly just quality of life
	payload = append(payload, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)

	return nil
}

// helper function to get ID from URL. This will be used extensivley throughout the application
func (app *application) readIDParam(r *http.Request) (int64, error) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil

}
