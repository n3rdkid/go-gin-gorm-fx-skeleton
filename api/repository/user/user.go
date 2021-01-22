package repository

import (
	"go-gin-gorm-fx-skeleton/models"
	"log"

	"gorm.io/gorm"
)

// UserRepository -> User Repository
type UserRepository interface {
	Save(models.User) (models.User, error)
	GetAll() ([]models.User, error)
	GetByID(string) (models.User, error)
}
type repository struct {
	DB *gorm.DB
}

// NewUserRepository -> returns new user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{DB: db}
}
func (u *repository) Migrate() error {
	log.Print("UserRepository :: Migrate")
	return u.DB.AutoMigrate(&models.User{})
}
func (u *repository) Save(user models.User) (models.User, error) {
	log.Print("UserRepository :: Save")
	err := u.DB.Create(&user).Error
	return user, err
}
func (u *repository) GetAll() (users []models.User, err error) {
	log.Print("UserRepository :: GetAll")
	return users, u.DB.Find(&users).Error
}
func (u *repository) GetByID(id string) (user models.User, err error) {
	log.Print("UserRepository :: GetByID")
	err = u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
