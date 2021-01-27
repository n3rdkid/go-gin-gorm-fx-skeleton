package routes

import (
	"go-gin-gorm-fx-skeleton/api/controller"
	"go-gin-gorm-fx-skeleton/lib"
)

// UserRoutes ->
type UserRoutes struct {
	handler        lib.RequestHandler
	userController controller.UserController
}

// Setup -> user routes
func (s UserRoutes) Setup() {
	api := s.handler.Gin.Group("/api")
	{
		api.GET("/user", s.userController.GetAll)
		api.GET("/user/:id", s.userController.GetByID)
		api.POST("/user", s.userController.Save)
	}
}

// NewUserRoutes -> Create new user controller
func NewUserRoutes(handler lib.RequestHandler, userController controller.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		userController: userController,
	}
}
