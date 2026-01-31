package user

import (
	domain "main/internal/domain/user"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service *domain.UserService
	logger  pkg.Logger
}

func NewUserController(service *domain.UserService, logger pkg.Logger) *UserController {
	return &UserController{
		service: service,
		logger:  logger,
	}
}

// Register godoc
//
// @Summary Register a new user
// @Description Register a new user account with email, password and optional username
// @Tags auth
// @Accept	json
// @Produce json
// @Param request body	UserRegisterRequest true "User registration request"
// @Success 200 {object} map[string]string	"Successful registration"
// @Failure 400 {object} map[string]string	"Invalid request or validation error"
// @Router	/api/v1/user/register [post]
func (controller *UserController) Register(ctx *gin.Context) {
	var req *UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err := controller.service.Register(req.Username, req.Password, req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user registered successfully",
	})

	controller.logger.Infow("user registered successfully",
		"username", req.Username,
	)
}

// Activate godoc
//
// @Summary Activate user account
// @Description Activate user account using verification token from email
// @Tags auth
// @Produce json
// @Param token query string true "Activation token from email"
// @Success 200 {object} map[string]string "User activated successfully"
// @Failure 400 {object} map[string]string "Token parameter is missing or invalid"
// @Router /api/v1/user/activate [get]
func (controller *UserController) Activate(ctx *gin.Context) {
	tokenStr, ok := ctx.GetQuery("token")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "token param is required"})

		return
	}

	if err := controller.service.Activate(tokenStr); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user activated successfully",
	})
}
