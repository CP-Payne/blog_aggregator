package feedfollows

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
	"github.com/CP-Payne/blog_aggregator/pkg/models"
	"github.com/google/uuid"
)

type FeedFollowHandler struct {
	util *helper.Util
	DB   *database.Queries
}

func NewFeedFollowsHandler(utils *helper.Util, db *database.Queries) *FeedFollowHandler {
	return &FeedFollowHandler{
		util: utils,
		DB:   db,
	}
}

// Follow a feed
func (h *FeedFollowHandler) FollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type paramaters struct {
		FeedId uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := paramaters{}
	err := decoder.Decode(&params)
	if err != nil {
		h.util.RespondWithError(w, http.StatusBadRequest, "Couldn't decode paramaters")
		return
	}
	// TODO: Check if feed exists before following
	feedFollow, err := h.DB.CreateFollowFeed(r.Context(), database.CreateFollowFeedParams{
		ID:        uuid.New(),
		FeedID:    params.FeedId,
		UserID:    user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		h.util.ServerError(w, err, "Couldn't follow feed")
		return
	}
	h.util.RespondWithJSON(w, http.StatusCreated, models.DatabaseFeedFollowToFeedFollow(feedFollow))
}

// Get all feeds a user follows
func (h *FeedFollowHandler) GetFeedsFollowByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	allUserFollowFeeds, err := h.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		h.util.ServerError(w, err, "Couldn't retrieve followed feeds")
		return
	}
	h.util.RespondWithJSON(w, http.StatusOK, models.DatabaseFeedsFollowToFeedsFollow(allUserFollowFeeds))
}

func (h *FeedFollowHandler) DeleteFeedsFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowID := r.PathValue("feedFollowID")
	uuidFeedFollowID, err := uuid.Parse(feedFollowID)
	if err != nil {
		h.util.RespondWithError(w, http.StatusBadRequest, "Invalid feed id provided")
		return
	}
	err = h.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     uuidFeedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		h.util.ServerError(w, err, "Something went wrong")
		return
	}
	type resp struct {
		Success string `json:"success"`
	}
	h.util.RespondWithJSON(w, http.StatusOK, resp{
		Success: "Feed unfollowed",
	})
}
