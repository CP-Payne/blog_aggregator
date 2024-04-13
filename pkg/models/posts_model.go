package models

import (
	"time"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Title       string    `json:"title"`
	Url         string    `json:"url"`
	Description string    `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	FeedID      uuid.UUID `json:"feed_id"`
}

func DatabasePostToPost(post database.GetPostsByUserRow) Post {
	return Post{
		ID:          post.ID,
		CreatedAt:   post.CreatedAt,
		UpdatedAt:   post.UpdatedAt,
		Title:       post.Title,
		Url:         post.Url,
		Description: post.Description.String,
		PublishedAt: post.PublishedAt.Time,
		FeedID:      post.FeedID,
	}
}

func DatabasePostsToPosts(posts []database.GetPostsByUserRow) []Post {
	result := make([]Post, len(posts))
	for i, post := range posts {
		result[i] = DatabasePostToPost(post)
	}
	return result
}
