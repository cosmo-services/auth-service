package bootstrap

import (
	"main/internal/application"
	"main/internal/config"
	"main/pkg"

	password_infrastructure "main/internal/infrastructure/password"
	user_infrastructure "main/internal/infrastructure/user"

	ping_application "main/internal/application/ping"
	swagger_application "main/internal/application/swagger"
	user_application "main/internal/application/user"

	password_domain "main/internal/domain/password"
	user_domain "main/internal/domain/user"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,

	password_infrastructure.Module,
	user_infrastructure.Module,

	application.Module,
	ping_application.Module,
	swagger_application.Module,
	user_application.Module,

	user_domain.Module,
	password_domain.Module,
)
