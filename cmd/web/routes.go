package main

import (
	"net/http"

	"github.com/CP-Payne/blog_aggregator/cmd/middleware"
	"github.com/CP-Payne/blog_aggregator/handlers/readiness"
	"github.com/CP-Payne/blog_aggregator/handlers/user"
)

func (app *application) routes() *http.ServeMux {
	middleware := middleware.NewMiddleware(app.util, app.DB)

	userHandler := user.NewUserHandler(app.util, app.DB)
	readinessHandler := readiness.NewReadinessHandler(app.util, app.DB)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readinessHandler.CheckReadiness)
	mux.HandleFunc("GET /v1/err", readinessHandler.CheckErr)

	// users

	mux.HandleFunc("POST /v1/users", userHandler.CreateUser)
	mux.HandleFunc("GET /v1/users", middleware.AuthMiddleware(userHandler.GetUser))

	return mux
}
