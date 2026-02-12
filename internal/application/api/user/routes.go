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
	api := r.handler.Gin.Group("/api/v2/auth/user")

	public := api.Group("/")
	{
		public.POST("/register", r.userController.Register)
		public.GET("/activate/confirm", r.userController.Activate)
	}

	protected := api.Group("/")
	protected.Use(r.authMiddleware.Handler())
	{
		protected.GET("/profile", r.userController.GetUser)
		protected.DELETE("/profile", r.userController.DeleteUser)
		protected.POST("/activate/resend", r.userController.ResendActivation)
		protected.POST("/email/change", r.userController.ChangeEmail)
		protected.POST("/password/change", r.userController.ChangePassword)
		protected.POST("/username/change", r.userController.ChangeUsername)
	}
}
