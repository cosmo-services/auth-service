package grpc_v1

import (
	"main/pkg"

	pb "github.com/cosmo-services/grpc-contracts/gen/api/v1"
	"go.uber.org/fx"
)

type GrpcHandler struct {
	grpc        pkg.GrpcServer
	authHandler *AuthHandler
}

func NewGrpcHandler(
	grpc pkg.GrpcServer,
	authHandler *AuthHandler,
) *GrpcHandler {
	return &GrpcHandler{
		grpc:        grpc,
		authHandler: authHandler,
	}
}

func (h *GrpcHandler) Setup() {
	pb.RegisterAuthServiceServer(
		h.grpc.Server,
		h.authHandler,
	)
}

var Module = fx.Options(
	fx.Provide(NewGrpcHandler),
	fx.Provide(NewAuthHandler),
)
