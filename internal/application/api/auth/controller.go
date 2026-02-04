package auth

import (
	"main/internal/domain/auth"
	"main/pkg"
	"net/http"

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
// @Router /login [post]
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

	ctx.JSON(http.StatusOK, pair)
}

// Refresh godoc
//
// @Summary	Refresh token
// @Description	Validate access token and return JWT tokens
// @Tags auth
// @Accept json
// @Produce	json
// @Param request body RefreshRequest true "Refresh credentials"
// @Success	200	{object} map[string]string "Successful authentication"
// @Failure	400	{object} map[string]string "Invalid request format"
// @Failure	401	{object} map[string]string "Invalid credentials"
// @Router /refresh [post]
func (controller *AuthController) Refresh(ctx *gin.Context) {
	var req *RefreshRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	pair, err := controller.authService.Refresh(req.RefreshToken)
	if err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, pair)
}
