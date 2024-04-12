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
	// When a user creates a feed, automatically follow it
	feedFollow, err := h.DB.CreateFollowFeed(r.Context(), database.CreateFollowFeedParams{
		ID:        uuid.New(),
		FeedID:    feed.ID,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		h.util.ServerError(w, err, "Something went wrong. Couldn't create follow feed")
		return
	}
	type resp struct {
		Feed       models.Feed       `json:"feed"`
		FeedFollow models.FeedFollow `json:"feed_follow"`
	}

	h.util.RespondWithJSON(w, http.StatusCreated, resp{
		Feed:       models.DatabaseFeedToFeed(feed),
		FeedFollow: models.DatabaseFeedFollowToFeedFollow(feedFollow),
	})
}

func (h *FeedHandler) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.DB.GetFeeds(r.Context())
	if err != nil {
		h.util.ServerError(w, err, "Couldn't retreive feeds")
		return
	}

	h.util.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedsToFeeds(feeds))
}
