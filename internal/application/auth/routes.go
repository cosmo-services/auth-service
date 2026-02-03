package auth

import (
	"main/pkg"
)

type AuthRoutes struct {
	handler        pkg.RequestHandler
	authController *AuthController
}

func NewAuthRoutes(
	authController *AuthController,
	handler pkg.RequestHandler,
) *AuthRoutes {
	return &AuthRoutes{
		authController: authController,
		handler:        handler,
	}
}

func (r *AuthRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/")
	{
		group.POST("/login", r.authController.Login)
		group.POST("/refresh", r.authController.Refresh)
	}
}
