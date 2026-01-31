package user

import (
	"main/pkg"
)

type UserRoutes struct {
	handler        pkg.RequestHandler
	userController *UserController
}

func NewUserRoutes(
	userController *UserController,
	handler pkg.RequestHandler,
) *UserRoutes {
	return &UserRoutes{
		userController: userController,
		handler:        handler,
	}
}

func (r *UserRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1/user")

	group.POST("/register", r.userController.Register)
	group.GET("/activate", r.userController.Activate)
}
