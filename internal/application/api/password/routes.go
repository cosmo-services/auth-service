package password_api

import (
	"main/pkg"
)

type PasswordRoutes struct {
	handler        pkg.RequestHandler
	pswdController *PasswordController
}

func NewPasswordRoutes(
	pswdController *PasswordController,
	handler pkg.RequestHandler,
) *PasswordRoutes {
	return &PasswordRoutes{
		pswdController: pswdController,
		handler:        handler,
	}
}

func (r *PasswordRoutes) Setup() {
	group := r.handler.Gin.Group("/api/v1")

	group.GET("/password/validate", r.pswdController.ValidatePassword)
}
