package main

import (
	"net/http"

	"github.com/CP-Payne/blog_aggregator/cmd/middleware"
	_ "github.com/lib/pq"
)

func main() {
	app := loadConfig()

	middleware := middleware.NewMiddleware(app.util, app.DB)
	corsMux := middleware.CorsMiddleware(app.routes())

	srv := &http.Server{
		Addr:     app.port,
		ErrorLog: app.util.ErrorLog,
		// Call the new app.routes method to get the servemux containing our routes
		Handler: corsMux,
	}

	app.util.InfoLog.Printf("Starting server on %s", app.port)
	// err := http.ListenAndServe(*addr, mux)
	err := srv.ListenAndServe()
	app.util.ErrorLog.Fatal(err)
}
