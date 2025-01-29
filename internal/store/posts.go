package store

import (
	"context"
	"database/sql"
	"github.com/hmprgm/swocial/models"
)

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context) error {

}