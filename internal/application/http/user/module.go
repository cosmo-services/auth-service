package user

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewUserRoutes),
)
