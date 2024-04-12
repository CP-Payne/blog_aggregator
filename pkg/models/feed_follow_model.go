package models

import (
	"time"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	FeedId    uuid.UUID `json:"feed_id"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func DatabaseFeedFollowToFeedFollow(feedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		FeedId:    feedFollow.FeedID,
		UserId:    feedFollow.UserID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func DatabaseFeedsFollowToFeedsFollow(feedsFollow []database.FeedFollow) []FeedFollow {
	result := make([]FeedFollow, len(feedsFollow))
	for i, feedFollow := range feedsFollow {
		result[i] = DatabaseFeedFollowToFeedFollow(feedFollow)
	}
	return result
}
