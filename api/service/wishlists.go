package service

import (
	"go-gin-gorm-fx-skeleton/api/repository"
	"go-gin-gorm-fx-skeleton/models"
)

// WishListsService ->
type WishListsService struct {
	wishlistRepository repository.WishListsRepository
}

// NewWishListsService -> returns new wishlist service
func NewWishListsService(repo repository.WishListsRepository) WishListsService {
	return WishListsService{
		wishlistRepository: repo,
	}
}

// Save ->
func (s *WishListsService) Save(wishlist models.WishList) (models.WishList, error) {
	return s.wishlistRepository.Save(wishlist)
}

// GetAll ->
func (s *WishListsService) GetAll() (wishlists []models.WishList, err error) {
	return s.wishlistRepository.GetAll()
}
