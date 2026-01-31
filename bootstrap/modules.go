package bootstrap

import (
	"main/internal/application"
	"main/internal/config"
	"main/pkg"

	auth_infrastructure "main/internal/infrastructure/auth"
	password_infrastructure "main/internal/infrastructure/password"
	tokens_infrastructure "main/internal/infrastructure/tokens"
	user_infrastructure "main/internal/infrastructure/user"

	auth_application "main/internal/application/auth"
	ping_application "main/internal/application/ping"
	swagger_application "main/internal/application/swagger"
	user_application "main/internal/application/user"

	auth_domain "main/internal/domain/auth"
	password_domain "main/internal/domain/password"
	tokens_domain "main/internal/domain/tokens"
	user_domain "main/internal/domain/user"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,

	password_infrastructure.Module,
	user_infrastructure.Module,
	auth_infrastructure.Module,
	tokens_infrastructure.Module,

	application.Module,
	ping_application.Module,
	swagger_application.Module,
	user_application.Module,
	auth_application.Module,

	user_domain.Module,
	password_domain.Module,
	auth_domain.Module,
	tokens_domain.Module,
)
