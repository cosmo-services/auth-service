package bootstrap

import (
	"main/internal/application/api"
	"main/internal/application/jobs"
	"main/internal/config"
	"main/internal/domain"
	"main/pkg"

	auth_infrastructure "main/internal/infrastructure/auth"
	password_infrastructure "main/internal/infrastructure/password"
	tokens_infrastructure "main/internal/infrastructure/tokens"
	user_infrastructure "main/internal/infrastructure/user"

	auth_api "main/internal/application/api/auth"
	health_api "main/internal/application/api/health"
	password_api "main/internal/application/api/password"
	swagger_api "main/internal/application/api/swagger"
	user_api "main/internal/application/api/user"

	auth_domain "main/internal/domain/auth"
	password_domain "main/internal/domain/password"
	tokens_domain "main/internal/domain/tokens"
	user_domain "main/internal/domain/user"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	config.Module,
	pkg.Module,
	domain.Module,

	password_infrastructure.Module,
	user_infrastructure.Module,
	auth_infrastructure.Module,
	tokens_infrastructure.Module,

	api.Module,
	jobs.Module,
	health_api.Module,
	swagger_api.Module,
	user_api.Module,
	auth_api.Module,
	password_api.Module,

	user_domain.Module,
	password_domain.Module,
	auth_domain.Module,
	tokens_domain.Module,
)
