package repository

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type Store struct {
	db     *sql.DB
	logger zerolog.Logger
}

func NewStore(db *sql.DB, logger zerolog.Logger) *Store {
	return &Store{
		db:     db,
		logger: logger,
	}
}
