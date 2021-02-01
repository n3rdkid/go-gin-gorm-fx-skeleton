package routes

import (
	"go-gin-gorm-fx-skeleton/api/controller"
	"go-gin-gorm-fx-skeleton/api/repository"
	"go-gin-gorm-fx-skeleton/lib"
	"log"
)

// WishlistsRoutes ->
type WishlistsRoutes struct {
	handler             lib.RequestHandler
	wishlistsController controller.WishListsController
	wishlistsRepository repository.WishListsRepository
}

// Setup -> wishlists routes
func (s WishlistsRoutes) Setup() {

	if err := s.wishlistsRepository.Migrate(); err != nil {
		log.Fatal("Wishlist migrate err", err)
	}
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/wishlists", s.wishlistsController.GetAll)
		api.POST("/wishlists", s.wishlistsController.Save)
	}
}

// NewWishListRoutes -> Create new wishlist controller
func NewWishListRoutes(handler lib.RequestHandler, wishlistsController controller.WishListsController, wishlistsRepository repository.WishListsRepository) WishlistsRoutes {
	return WishlistsRoutes{
		handler:             handler,
		wishlistsController: wishlistsController,
		wishlistsRepository: wishlistsRepository,
	}
}
