package repository

import (
	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/models"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user repository operations.
type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uuid.UUID) (*models.User, error)
	Update(user *models.User) error
	Delete(id uuid.UUID) error
	List() ([]*models.User, error)
}

// VideoRepository defines the interface for video repository operations.
type VideoRepository interface {
	Create(video *models.Video) error
	GetByID(id uuid.UUID) (*models.Video, error)
	Update(video *models.Video) error
	Delete(id uuid.UUID) error
	List() ([]*models.Video, error)
}

// Repository struct holds instances of user and video repositories.
type Repository struct {
	UserRepo  UserRepository
	VideoRepo VideoRepository
}

// NewRepository creates a new instance of Repository.
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo:  NewUserRepository(db),
		VideoRepo: NewVideoRepository(db),
	}
}
