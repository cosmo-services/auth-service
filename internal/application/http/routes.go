package http

import (
	auth_http "main/internal/application/http/auth"
	health_http "main/internal/application/http/health"
	password_http "main/internal/application/http/password"
	swagger_http "main/internal/application/http/swagger"
	user_http "main/internal/application/http/user"

	"go.uber.org/fx"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swaggerRoutes *swagger_http.SwaggerRoutes,
	healthRoutes *health_http.HealthRoutes,
	userRoutes *user_http.UserRoutes,
	auhtRoutes *auth_http.AuthRoutes,
	pswdRoutes *password_http.PasswordRoutes,
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
