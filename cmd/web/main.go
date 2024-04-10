package main

import "net/http"

func main() {
	app := loadConfig()

	srv := &http.Server{
		Addr:     app.port,
		ErrorLog: app.errorLog,
		// Call the new app.routes method to get the servemux containing our routes
		Handler: app.routes(),
	}

	app.infoLog.Printf("Starting server on %s", app.port)
	// err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}
