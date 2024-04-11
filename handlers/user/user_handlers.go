package user

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
	"github.com/CP-Payne/blog_aggregator/pkg/models"
	"github.com/google/uuid"
)

type UserHandler struct {
	util *helper.Util
	DB   *database.Queries
}

func NewUserHandler(utils *helper.Util, db *database.Queries) *UserHandler {
	return &UserHandler{
		util: utils,
		DB:   db,
	}
}

// Create user

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	type params struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	userParams := params{}
	err := decoder.Decode(&userParams)
	if err != nil {
		h.util.RespondWithError(w, http.StatusInternalServerError, "Couldn't decode paramters")
		return
	}

	user, err := h.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userParams.Name,
	})
	if err != nil {
		h.util.ServerError(w, err, "Couldn't create user")
		return
	}
	h.util.RespondWithJSON(w, http.StatusCreated, models.DatabaseUserToUser(user))
}

// Get user by ApiKey
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	h.util.RespondWithJSON(w, http.StatusOK, models.DatabaseUserToUser(user))
}
