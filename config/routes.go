package config

import (
	"go-gin-gorm-fx-skeleton/api/routes"
	"go-gin-gorm-fx-skeleton/utils"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes -> Setup routes
func InitializeRoutes(db *gorm.DB) {
	httpRouter := gin.Default()
	httpRouter.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	routes.UserRoutes(httpRouter.Group("user"), db)
	port := utils.GetEnvByKey("SERVER_PORT")
	if port == "" {
		httpRouter.Run()
	} else {
		httpRouter.Run(os.Getenv("SERVER_PORT"))
	}
}
