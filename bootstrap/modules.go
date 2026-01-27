package bootstrap

import (
	"main/internal/api"
	"main/internal/infrastructure"

	"main/internal/config"
	"main/pkg"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	api.Module,
	infrastructure.Module,
)
