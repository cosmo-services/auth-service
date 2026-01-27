package api

import (
	ping_api "main/internal/api/ping"
	swagger_api "main/internal/api/swagger"
)

type Route interface {
	Setup()
}

type Routes []Route

func NewRoutes(
	swaggerRoutes *swagger_api.SwaggerRoutes,
	pingRoutes *ping_api.PingRoutes,
) Routes {
	return Routes{
		pingRoutes,
		swaggerRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
