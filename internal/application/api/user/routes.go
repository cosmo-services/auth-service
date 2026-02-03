package user

import (
	"main/internal/application/api/auth"
	"main/pkg"
)

type UserRoutes struct {
	handler        pkg.RequestHandler
	userController *UserController
	authMiddleware *auth.AuthMiddleware
}

func NewUserRoutes(
	userController *UserController,
	handler pkg.RequestHandler,
	authMiddleware *auth.AuthMiddleware,
) *UserRoutes {
	return &UserRoutes{
		userController: userController,
		handler:        handler,
		authMiddleware: authMiddleware,
	}
}

func (r *UserRoutes) Setup() {
	api := r.handler.Gin.Group("/api/v1")

	public := api.Group("/")
	{
		public.POST("/register", r.userController.Register)
		public.GET("/activate", r.userController.Activate)
	}

	protected := api.Group("/")
	protected.Use(r.authMiddleware.Handler())
	{
		protected.POST("/activate/resend", r.userController.ResendActivation)
	}
}
