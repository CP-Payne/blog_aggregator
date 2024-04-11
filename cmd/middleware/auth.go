package middleware

import (
	"net/http"

	"github.com/CP-Payne/blog_aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (mw *middleware) AuthMiddleware(handler authHandler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apikey, err := mw.util.GetApiKey(r.Header)
		if err != nil {
			mw.util.RespondWithError(w, http.StatusUnauthorized, "Invalid ApiKey")
			return
		}

		user, err := mw.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			mw.util.RespondWithError(w, http.StatusUnauthorized, "Couldn't retrieve user")
			return
		}
		handler(w, r, user)
	})
}
