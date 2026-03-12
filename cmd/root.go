package cmd

import (
	"fmt"
	"log"
	"main/bootstrap"
	"net"

	"context"

	"go.uber.org/fx"

	"main/pkg"

	grpc_v1 "main/internal/application/grpc/v1"
	"main/internal/application/http"
	"main/internal/application/jobs"
	"main/internal/application/nats"
	"main/internal/config"
)

func SetupApp(
	lc fx.Lifecycle,
	env config.Env,
	logger pkg.Logger,
	handler pkg.RequestHandler,
	routes http.Routes,
	nats *nats.Nats,
	grpc pkg.GrpcServer,
	grpcHandler *grpc_v1.GrpcHandler,
	workers jobs.Workers,
) {
	ctx, cancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			go func() {
				nats.SetupPublishers()
			}()

			go func() {
				grpcHandler.Setup()

				lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.GrcpPort))
				if err != nil {
					log.Fatal(err)
				}

				if err := grpc.Server.Serve(lis); err != nil {
					log.Fatal(err)
				}
			}()

			go func() {
				workers.Run(ctx)
			}()

			go func() {
				routes.Setup()

				if err := handler.Gin.Run(":" + env.Port); err != nil {
					logger.Error(err)
				}
			}()

			return nil
		},
		OnStop: func(stopCtx context.Context) error {
			cancel()

			return nil
		},
	})
}

func StartApp() {
	opts := fx.Options(
		fx.Invoke(SetupApp),
	)

	app := fx.New(
		bootstrap.CommonModules,
		opts,
	)

	app.Run()
}
