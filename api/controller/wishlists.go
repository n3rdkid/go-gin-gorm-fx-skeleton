package controller

import (
	"go-gin-gorm-fx-skeleton/api/service"
	"go-gin-gorm-fx-skeleton/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WishListsController ->
type WishListsController struct {
	service service.WishListsService
}

// NewWishListsController -> NewWishListsController constructor
func NewWishListsController(wishlistService service.WishListsService) WishListsController {
	return WishListsController{
		service: wishlistService,
	}
}

// Save ->
func (controller *WishListsController) Save(c *gin.Context) {
	log.Println("Inside Save Wishlist")
	var wishlist models.WishList
	if err := c.ShouldBindJSON(&wishlist); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wishlist, err1 := controller.service.Save(wishlist)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": wishlist})
}

// GetAll ->
func (controller *WishListsController) GetAll(c *gin.Context) {
	log.Println("Inside GetAll WishLists")
	wishlists, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": wishlists,
	})
}
