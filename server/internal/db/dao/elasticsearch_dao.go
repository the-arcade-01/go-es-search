package dao

import "github.com/the-arcade-01/go-es-search/server/internal/models/db"

type ElasticSearchDao interface {
	Search() ([]*db.ProductDB, error)
}
