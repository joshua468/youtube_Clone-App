package app

import (
	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/repository"
	"github.com/joshua468/youtube_Clone/utils"
)

type UserApp struct {
	Repo *repository.UserRepository
}

func NewUserApp(repo *repository.UserRepository) *UserApp {
	return &UserApp{
		Repo: repo,
	}
}

func (ua *UserApp) GetUserByID(userID uuid.UUID) (*models.User, error) {
	user, err := ua.Repo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ua *UserApp) CreateUser(req *models.CreateUserRequest) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		IsAdmin:  req.IsAdmin,
	}

	err = ua.Repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
