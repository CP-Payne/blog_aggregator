package main

import (
	"net/http"

	"github.com/CP-Payne/blog_aggregator/cmd/middleware"
	"github.com/CP-Payne/blog_aggregator/handlers/feed"
	feedfollows "github.com/CP-Payne/blog_aggregator/handlers/feed_follows"
	"github.com/CP-Payne/blog_aggregator/handlers/readiness"
	"github.com/CP-Payne/blog_aggregator/handlers/user"
)

func (app *application) routes() *http.ServeMux {
	// Middleware defined here are applied on all requests and responses
	// Middleware in the routes function are applied on specific endpoints
	middleware := middleware.NewMiddleware(app.util, app.DB)

	readinessHandler := readiness.NewReadinessHandler(app.util, app.DB)
	userHandler := user.NewUserHandler(app.util, app.DB)
	feedHandler := feed.NewFeedHandler(app.util, app.DB)
	feedFollowHandler := feedfollows.NewFeedFollowsHandler(app.util, app.DB)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/readiness", readinessHandler.CheckReadiness)
	mux.HandleFunc("GET /v1/err", readinessHandler.CheckErr)

	// users
	mux.HandleFunc("POST /v1/users", userHandler.CreateUser)
	mux.HandleFunc("GET /v1/users", middleware.AuthMiddleware(userHandler.GetUser))

	// feeds
	mux.HandleFunc("POST /v1/feeds", middleware.AuthMiddleware(feedHandler.CreateFeed))
	mux.HandleFunc("GET /v1/feeds", feedHandler.GetFeeds)

	// Following Feeds
	mux.HandleFunc("POST /v1/feed_follows", middleware.AuthMiddleware(feedFollowHandler.FollowFeed))
	mux.HandleFunc("GET /v1/feed_follows", middleware.AuthMiddleware(feedFollowHandler.GetFeedsFollowByUser))
	mux.HandleFunc("DELETE /v1/feed_follows/{feedFollowID}", middleware.AuthMiddleware(feedFollowHandler.DeleteFeedsFollow))

	return mux
}
