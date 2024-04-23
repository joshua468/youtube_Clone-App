package services

import (
	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/repository"
)

type VideoServiceImpl struct {
	VideoRepo repository.VideoRepository
}

func NewVideoService(videoRepo repository.VideoRepository) VideoService {
	return &VideoServiceImpl{
		VideoRepo: videoRepo,
	}
}

func (s *VideoServiceImpl) CreateVideo(video *models.Video) error {
	return s.VideoRepo.Create(video)
}

func (s *VideoServiceImpl) GetVideoByID(id string) (*models.Video, error) {
	videoID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	return s.VideoRepo.GetByID(videoID)
}

func (s *VideoServiceImpl) UpdateVideo(id string, video *models.Video) error {
	videoID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	existingVideo, err := s.VideoRepo.GetByID(videoID)
	if err != nil {
		return err
	}

	existingVideo.Title = video.Title
	existingVideo.Description = video.Description
	existingVideo.URL = video.URL
	existingVideo.UserID = video.UserID

	return s.VideoRepo.Update(existingVideo)
}

func (s *VideoServiceImpl) DeleteVideo(id string) error {
	videoID, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return s.VideoRepo.Delete(videoID)
}

func (s *VideoServiceImpl) ListVideos(page, pageSize int) ([]*models.Video, error) {
	return s.VideoRepo.List()
}
