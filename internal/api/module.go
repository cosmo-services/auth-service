package api

import (
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewRoutes),

	ping_api.Module,
	swagger_api.Module,
)
