package cmd

import (
	"main/bootstrap"

	"context"

	"go.uber.org/fx"

	"main/pkg"

	"main/internal/application/api"
	"main/internal/application/jobs"
	"main/internal/config"
)

func Run() any {
	return func(
		env config.Env,
		logger pkg.Logger,
		handler pkg.RequestHandler,
		routes api.Routes,
		workers jobs.Workers,
	) {
		routes.Setup()
		workers.Run()

		err := handler.Gin.Run(":" + env.Port)

		if err != nil {
			logger.Error(err)
			return
		}
	}
}

func StartApp() error {
	opts := fx.Options(
		fx.Invoke(Run()),
	)

	app := fx.New(
		bootstrap.CommonModules,
		opts,
	)
	ctx := context.Background()
	err := app.Start(ctx)
	return err
}
