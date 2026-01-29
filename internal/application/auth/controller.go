package auth

import (
	"main/internal/domain/auth"
	"main/pkg"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *auth.AuthService
	logger      pkg.Logger
}

func NewAuthController(
	authService *auth.AuthService,
	logger pkg.Logger,
) *AuthController {
	return &AuthController{
		authService: authService,
		logger:      logger,
	}
}

// Login godoc
//
// @Summary	User login
// @Description	Authenticate user and return JWT tokens
// @Tags auth
// @Accept json
// @Produce	json
// @Param request body LoginRequest	true "Login credentials"
// @Success	200	{object} map[string]string "Successful authentication"
// @Failure	400	{object} map[string]string "Invalid request format"
// @Failure	401	{object} map[string]string "Invalid credentials"
// @Router /api/v1/auth/login [post]
func (controller *AuthController) Login(ctx *gin.Context) {
	var req *LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	pair, err := controller.authService.Login(req.Username, req.Password)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token_type":    "Bearer",
		"access_token":  pair.Access.Token,
		"refresh_token": pair.Refresh.Token,
		"expires_in":    int(time.Until(pair.Access.Expires).Seconds()),
	})
}
