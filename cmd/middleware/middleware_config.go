package middleware

import (
	"github.com/CP-Payne/blog_aggregator/internal/database"
	"github.com/CP-Payne/blog_aggregator/pkg/helper"
)

type middleware struct {
	util *helper.Util
	DB   *database.Queries
}

func NewMiddleware(util *helper.Util, db *database.Queries) *middleware {
	return &middleware{
		util: util,
		DB:   db,
	}
}
