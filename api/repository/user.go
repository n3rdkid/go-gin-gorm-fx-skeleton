package repository

import (
	"go-gin-gorm-fx-skeleton/lib"
	"go-gin-gorm-fx-skeleton/models"
	"log"
)

// UserRepository ->
type UserRepository struct {
	db lib.Database
}

// NewUserRepository -> returns new user repository
func NewUserRepository(db lib.Database) UserRepository {
	return UserRepository{db: db}
}

// Save ->
func (u *UserRepository) Save(user models.User) (models.User, error) {
	log.Print("UserRepository :: Save")
	err := u.db.DB.Create(&user).Error
	return user, err
}

// GetAll ->
func (u *UserRepository) GetAll() (users []models.User, err error) {
	log.Print("UserRepository :: GetAll")
	log.Println(users)
	return users, u.db.DB.Find(&users).Error
}

// GetByID ->
func (u *UserRepository) GetByID(id string) (user models.User, err error) {
	log.Print("UserRepository :: GetByID")
	err = u.db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
