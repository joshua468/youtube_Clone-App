package app

import (
	"github.com/joshua468/youtube_Clone/repository"
	"github.com/rs/zerolog"
)

type App struct {
	UserRepo  repository.UserRepository
	VideoRepo repository.VideoRepository
	Logger    zerolog.Logger
}

func NewApp(userRepo repository.UserRepository, videoRepo repository.VideoRepository, logger zerolog.Logger) *App {
	return &App{
		UserRepo:  userRepo,
		VideoRepo: videoRepo,
		Logger:    logger,
	}
}
