package bootstrap

import (
	"main/internal/application/http"
	"main/internal/application/jobs"
	"main/internal/config"
	"main/internal/domain"
	"main/pkg"

	auth_infrastructure "main/internal/infrastructure/auth"
	password_infrastructure "main/internal/infrastructure/password"
	tokens_infrastructure "main/internal/infrastructure/tokens"
	user_infrastructure "main/internal/infrastructure/user"

	auth_http "main/internal/application/http/auth"
	health_http "main/internal/application/http/health"
	password_http "main/internal/application/http/password"
	swagger_http "main/internal/application/http/swagger"
	user_http "main/internal/application/http/user"

	nats "main/internal/application/nats"

	grpc_v1 "main/internal/application/grpc/v1"

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

	http.Module,
	jobs.Module,
	nats.Module,
	grpc_v1.Module,
	health_http.Module,
	swagger_http.Module,
	user_http.Module,
	auth_http.Module,
	password_http.Module,

	user_domain.Module,
	password_domain.Module,
	auth_domain.Module,
	tokens_domain.Module,
)
