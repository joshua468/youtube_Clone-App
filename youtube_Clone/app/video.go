package app

import (
	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/repository"
)

type VideoApp struct {
	Repo *repository.VideoRepository
}

func NewVideoApp(repo *repository.VideoRepository) *VideoApp {
	return &VideoApp{
		Repo: repo,
	}
}

func (va *VideoApp) GetVideoByID(videoID uuid.UUID) (*models.Video, error) {
	video, err := va.Repo.GetByID(videoID)
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (va *VideoApp) CreateVideo(req *models.CreateVideoRequest) (*models.Video, error) {
	video := &models.Video{
		Title:       req.Title,
		Description: req.Description,
		URL:         req.URL,
		UserID:      req.UserID,
	}

	err := va.Repo.Create(video)
	if err != nil {
		return nil, err
	}
	return video, nil
}
