package tokens

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewTokenRepository),
	fx.Provide(NewSha256Hasher),
	fx.Provide(NewBase64TokenGenerator),
)
