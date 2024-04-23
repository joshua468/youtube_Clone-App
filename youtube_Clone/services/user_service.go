package services

import (
	"github.com/google/uuid"
	"github.com/joshua468/youtube_Clone/models"
	"github.com/joshua468/youtube_Clone/repository"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
	}
}

func (s *UserServiceImpl) CreateUser(user *models.User) error {
	return s.UserRepo.Create(user)
}

func (s *UserServiceImpl) GetUserByID(id string) (*models.User, error) {

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return s.UserRepo.GetByID(userID)
}

// UpdateUser updates an existing user.
func (s *UserServiceImpl) UpdateUser(id string, user *models.User) error {
	// Convert ID string to UUID.
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	existingUser, err := s.UserRepo.GetByID(userID)
	if err != nil {
		return err
	}

	// Update fields of existingUser based on the input user.
	existingUser.Username = user.Username
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.IsAdmin = user.IsAdmin

	return s.UserRepo.Update(existingUser)
}

// DeleteUser deletes a user by ID.
func (s *UserServiceImpl) DeleteUser(id string) error {
	// Convert ID string to UUID.
	userID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return s.UserRepo.Delete(userID)
}

// ListUsers retrieves a list of all users.
func (s *UserServiceImpl) ListUsers(offset, pageSize int) ([]*models.User, error) {
	return s.UserRepo.List()
}
