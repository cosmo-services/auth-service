package password_api

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewPasswordController),
	fx.Provide(NewPasswordRoutes),
)
