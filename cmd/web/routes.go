package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", app.checkReadiness)
	mux.HandleFunc("GET /v1/err", app.checkErr)

	return mux
}
