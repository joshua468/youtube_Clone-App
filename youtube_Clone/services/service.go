package services

import (
	"github.com/joshua468/youtube_Clone/models"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	UpdateUser(id string, user *models.User) error
	DeleteUser(id string) error
	ListUsers(offset, pageSize int) ([]*models.User, error) // Updated signature
}

type VideoService interface {
	CreateVideo(video *models.Video) error
	GetVideoByID(id string) (*models.Video, error)
	UpdateVideo(id string, video *models.Video) error
	DeleteVideo(id string) error
	ListVideos(page, pageSize int) ([]*models.Video, error) // Update signature
}
