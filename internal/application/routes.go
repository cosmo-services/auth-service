package application

import (
	auth_application "main/internal/application/auth"
	ping_application "main/internal/application/ping"
	swagger_application "main/internal/application/swagger"
	user_application "main/internal/application/user"

	"go.uber.org/fx"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swaggerRoutes *swagger_application.SwaggerRoutes,
	pingRoutes *ping_application.PingRoutes,
	userRoutes *user_application.UserRoutes,
	auhtRoutes *auth_application.AuthRoutes,
) Routes {
	return Routes{
		pingRoutes,
		swaggerRoutes,
		userRoutes,
		auhtRoutes,
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
