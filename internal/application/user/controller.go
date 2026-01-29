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
