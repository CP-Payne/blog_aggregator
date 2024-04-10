package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	// In addition to informing th user about the internal server error
	// We also want to know what the internal server error is in order to fix it.
	// Therefore, we also print it to stdOut
	app.errorLog.Output(2, trace)

	// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	app.respondWithError(w, 500, http.StatusText(http.StatusInternalServerError))
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) respondWithError(w http.ResponseWriter, code int, message string) {
	// if code > 499 {
	// 	log.Printf("Responing with 5xx error: %s", message)
	// }
	type errorResponse struct {
		Error string `json:"error"`
	}
	app.respondWithJSON(w, code, errorResponse{
		Error: message,
	})
}

func (app *application) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}
