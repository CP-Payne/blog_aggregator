package main

import (
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	app := loadConfig()

	srv := &http.Server{
		Addr:     app.port,
		ErrorLog: app.util.ErrorLog,
		// Call the new app.routes method to get the servemux containing our routes
		Handler: app.routes(),
	}

	app.util.InfoLog.Printf("Starting server on %s", app.port)
	// err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()
	app.util.ErrorLog.Fatal(err)
}
