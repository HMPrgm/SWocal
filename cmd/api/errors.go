package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s %s path: %s, error: %s", "internal server error", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s %s path: %s, error: %s", "Bad Request", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusBadRequest, err.Error())
}
func (app *application) notFound(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("%s %s path: %s, error: %s", "Not Found", r.Method, r.URL.Path, err)

	writeJSONError(w, http.StatusNotFound, "Not Found")
}
