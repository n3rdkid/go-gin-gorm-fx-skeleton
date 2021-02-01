package service

import (
	"go-gin-gorm-fx-skeleton/api/repository"
	"go-gin-gorm-fx-skeleton/models"
)

// UserService ->
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(repo repository.UserRepository) UserService {
	return UserService{
		userRepository: repo,
	}
}

// Save ->
func (s *UserService) Save(user models.User) (models.User, error) {
	_, err := s.userRepository.GetByID(user.ID)
	if err == nil {
		return user, err
	}
	return s.userRepository.Save(user)
}

// GetAll ->
func (s *UserService) GetAll() (users []models.User, err error) {
	return s.userRepository.GetAll()
}

// GetByID ->
func (s *UserService) GetByID(id string) (models.User, error) {
	return s.userRepository.GetByID(id)
}