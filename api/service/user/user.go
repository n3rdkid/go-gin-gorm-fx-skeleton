package service

import (
	repository "go-gin-gorm-fx-skeleton/api/repository/user"
	"go-gin-gorm-fx-skeleton/models"
)

// UserService -> User Service Contract
type UserService interface {
	Save(models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetByID(string) (models.User, error)
}

type service struct {
	userRepository repository.UserRepository
}

// NewUserService -> returns new user service
func NewUserService(repo repository.UserRepository) UserService {
	return &service{
		userRepository: repo,
	}
}
func (s *service) Save(user models.User) (models.User, error) {
	_, err := s.userRepository.GetByID(user.ID)
	if err == nil {
		return user, err
	}
	return s.userRepository.Save(user)
}

func (s *service) GetAll() (users []models.User, err error) {
	return s.userRepository.GetAll()
}

func (s *service) GetByID(id string) (models.User, error) {
	return s.userRepository.GetByID(id)
}
