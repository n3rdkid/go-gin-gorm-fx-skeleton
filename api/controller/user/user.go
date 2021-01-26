package controller

import (
	service "go-gin-gorm-fx-skeleton/api/service/user"
	"go-gin-gorm-fx-skeleton/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController -> User Controller Contract
type UserController interface {
	Save(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type controller struct {
	service service.UserService
}

// NewUserController -> NewUserController constructor
func NewUserController(s service.UserService) UserController {
	return &controller{
		service: s,
	}
}
func (controller *controller) Save(c *gin.Context) {
	log.Println("Inside Save User")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	controller.service.Save(user)
	c.JSON(http.StatusCreated, gin.H{"data": user})
}
func (controller *controller) GetAll(c *gin.Context) {
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
func (controller *controller) GetByID(c *gin.Context) {
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
