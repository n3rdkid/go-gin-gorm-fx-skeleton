package routes

import (
	controller "go-gin-gorm-fx-skeleton/api/controller/user"
	repository "go-gin-gorm-fx-skeleton/api/repository/user"
	service "go-gin-gorm-fx-skeleton/api/service/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRoutes -> UserRoutes
func UserRoutes(route *gin.RouterGroup, db *gorm.DB) {

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	if err := userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	route.GET("/", userController.GetAll)
	route.GET("/:id", userController.GetByID)
	route.POST("/", userController.Save)

}
