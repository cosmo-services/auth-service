package api

import (
	auth_api "main/internal/application/api/auth"
	health_api "main/internal/application/api/health"
	password_api "main/internal/application/api/password"
	swagger_api "main/internal/application/api/swagger"
	user_api "main/internal/application/api/user"

	"go.uber.org/fx"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swaggerRoutes *swagger_api.SwaggerRoutes,
	healthRoutes *health_api.HealthRoutes,
	userRoutes *user_api.UserRoutes,
	auhtRoutes *auth_api.AuthRoutes,
	pswdRoutes *password_api.PasswordRoutes,
) Routes {
	return Routes{
		healthRoutes,
		swaggerRoutes,
		userRoutes,
		auhtRoutes,
		pswdRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}

var Module = fx.Options(
	fx.Provide(NewRoutes),
)
