package readiness

import (
	"errors"
	"net/http"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
)

type ReadinessHandler struct {
	util *helper.Util
	DB   *database.Queries
}

func NewReadinessHandler(utils *helper.Util, db *database.Queries) *ReadinessHandler {
	return &ReadinessHandler{
		util: utils,
		DB:   db,
	}
}

// Testing helper functions
func (h *ReadinessHandler) CheckReadiness(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}

	h.util.RespondWithJSON(w, 200, Response{
		Status: "ok",
	})
}

// Testing helper functions
func (h *ReadinessHandler) CheckErr(w http.ResponseWriter, r *http.Request) {
	type Response struct {
		Status string `json:"status"`
	}

	h.util.ServerError(w, errors.New("something went wrong"), "something went wrong")
}
