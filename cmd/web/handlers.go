package main

import (
	"errors"
	"net/http"
)

// Testing helper functions
func (app *application) checkReadiness(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}

	app.respondWithJSON(w, 200, Response{
		Status: "ok",
	})
}

// Testing helper functions
func (app *application) checkErr(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}

	app.serverError(w, errors.New("something went wrong"))
}
