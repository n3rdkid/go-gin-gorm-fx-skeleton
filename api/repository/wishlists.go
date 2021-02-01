package repository

import (
	"go-gin-gorm-fx-skeleton/lib"
	"go-gin-gorm-fx-skeleton/models"
	"log"
)

// WishListsRepository ->
type WishListsRepository struct {
	db lib.Database
}

// NewWishListsRepository -> returns new wishlists repository
func NewWishListsRepository(db lib.Database) WishListsRepository {
	return WishListsRepository{db: db}
}

// Migrate ->
func (u *WishListsRepository) Migrate() error {
	log.Print("[WishListsRepository]...Migrate")
	return u.db.DB.AutoMigrate(&models.WishList{})
}

// Save ->
func (u *WishListsRepository) Save(wishlist models.WishList) (models.WishList, error) {
	log.Print("WishListsRepository :: Save")
	err := u.db.DB.Create(&wishlist).Error
	return wishlist, err
}

// GetAll ->
func (u *WishListsRepository) GetAll() (wishlists []models.WishList, err error) {
	log.Print("WishListsRepository :: GetAll")
	return wishlists, u.db.DB.Find(&wishlists).Error
}
