package feed

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
	"github.com/CP-Payne/blog_aggregator/pkg/models"
	"github.com/google/uuid"
)

type FeedHandler struct {
	util *helper.Util
	DB   *database.Queries
}

func NewFeedHandler(utils *helper.Util, db *database.Queries) *FeedHandler {
	return &FeedHandler{
		util: utils,
		DB:   db,
	}
}

func (h *FeedHandler) CreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramaters{}
	err := decoder.Decode(&params)
	if err != nil {
		h.util.RespondWithError(w, http.StatusBadRequest, "Couldn't decode parameters")
		return
	}

	feed, err := h.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		h.util.ServerError(w, err, "Couldn't create feed")
		return
	}
	h.util.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedToFeed(feed))
}
