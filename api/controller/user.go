package controller

import (
	"go-gin-gorm-fx-skeleton/api/service"
	"go-gin-gorm-fx-skeleton/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController ->
type UserController struct {
	service service.UserService
}

// NewUserController -> NewUserController constructor
func NewUserController(userService service.UserService) UserController {
	return UserController{
		service: userService,
	}
}

// Save ->
func (controller *UserController) Save(c *gin.Context) {
	log.Println("Inside Save User")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.service.Save(user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// GetAll ->
func (controller *UserController) GetAll(c *gin.Context) {
	log.Println("Inside GetAll users")
	users, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

// GetByID ->
func (controller *UserController) GetByID(c *gin.Context) {
	log.Println("Inside GetByID")
	id := c.Param("id")
	user, err := controller.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
