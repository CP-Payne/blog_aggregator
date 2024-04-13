package posts

import (
	"net/http"
	"strconv"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
	"github.com/CP-Payne/blog_aggregator/pkg/models"
)

type PostsHandler struct {
	util *helper.Util
	DB   *database.Queries
}

func NewPostsHandler(utils *helper.Util, db *database.Queries) *PostsHandler {
	return &PostsHandler{
		util: utils,
		DB:   db,
	}
}

func (h *PostsHandler) GetPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) {
	queries := r.URL.Query()
	limitParam := queries.Get("limit")

	limitInt := 10

	if limitParam != "" {
		var err error
		limitInt, err = strconv.Atoi(limitParam)
		if err != nil {
			h.util.RespondWithError(w, http.StatusBadRequest, "Invalid limit paramater provided")
			return
		}
	}

	// userPosts is the will contain the latests posts from the rss feeds the user is subscribed to
	userPosts, err := h.DB.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		UserID: user.ID,
		Limit:  int32(limitInt),
	})
	if err != nil {
		h.util.ServerError(w, err, "Couldn't retrieve latest posts")
		return
	}

	h.util.RespondWithJSON(w, http.StatusOK, models.DatabasePostsToPosts(userPosts))
}
