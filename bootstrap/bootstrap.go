package bootstrap

import (
	"context"
	"go-gin-gorm-fx-skeleton/api/controller"
	"go-gin-gorm-fx-skeleton/api/repository"
	"go-gin-gorm-fx-skeleton/api/routes"
	"go-gin-gorm-fx-skeleton/api/service"
	"go-gin-gorm-fx-skeleton/lib"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	lib.Module,
	controller.Module,
	routes.Module,
	service.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler lib.RequestHandler,
	routes routes.Routes,
	env lib.Env,
	database lib.Database,
) {
	conn, _ := database.DB.DB()

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			conn.SetMaxOpenConns(10)
			go func() {
				routes.Setup()
				handler.Gin.Run(env.ServerPort)
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			conn.Close()
			return nil
		},
	})
}
