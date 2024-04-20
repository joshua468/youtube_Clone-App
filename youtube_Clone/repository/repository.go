package repository

import (
	"database/sql"

	"github.com/rs/zerolog"
)

type Repository struct {
	UserRepo  *UserRepository
	VideoRepo *VideoRepository
}

func NewRepository(db *sql.DB, logger zerolog.Logger) *Repository {
	return &Repository{
		UserRepo:  NewUserRepository(db),
		VideoRepo: NewVideoRepository(db),
	}
}
