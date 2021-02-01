package routes

import (
	"go-gin-gorm-fx-skeleton/api/controller"
	"go-gin-gorm-fx-skeleton/api/repository"
	"go-gin-gorm-fx-skeleton/lib"
	"log"
)

// UserRoutes ->
type UserRoutes struct {
	handler        lib.RequestHandler
	userController controller.UserController
	userRepository repository.UserRepository
}

// Setup -> user routes
func (s UserRoutes) Setup() {

	if err := s.userRepository.Migrate(); err != nil {
		log.Fatal("User migrate err", err)
	}
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetAll)
		api.GET("/user/:id", s.userController.GetByID)
		api.POST("/user", s.userController.Save)
	}
}

// NewUserRoutes -> Create new user controller
func NewUserRoutes(handler lib.RequestHandler, userController controller.UserController, userRepository repository.UserRepository) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
		userRepository: userRepository,
	}
}
