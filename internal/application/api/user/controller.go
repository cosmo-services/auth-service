package user

import (
	"errors"
	domain "main/internal/domain/user"
	"main/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *domain.UserService
	logger      pkg.Logger
}

func NewUserController(service *domain.UserService, logger pkg.Logger) *UserController {
	return &UserController{
		userService: service,
		logger:      logger,
	}
}

// Register godoc
//
// @Summary Register a new user
// @Description Register a new user account with email, password and optional username
// @Tags user
// @Accept	json
// @Produce json
// @Param request body	UserRegisterRequest true "User registration request"
// @Success 200 {object} map[string]string	"Successful registration"
// @Failure 400 {object} map[string]string	"Invalid request or validation error"
// @Router	/user/register [post]
func (controller *UserController) Register(ctx *gin.Context) {
	var req *UserRegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	err := controller.userService.Register(req.Username, req.Password, req.Email)
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
// @Tags user
// @Produce json
// @Param token query string true "Activation token from email"
// @Success 200 {object} map[string]string "User activated successfully"
// @Failure 400 {object} map[string]string "Token parameter is missing or invalid"
// @Router /user/activate/confirm [get]
func (controller *UserController) Activate(ctx *gin.Context) {
	tokenStr, ok := ctx.GetQuery("token")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "token param is required"})

		return
	}

	if err := controller.userService.Activate(tokenStr); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user activated successfully",
	})
}

// ResendActivation godoc
//
// @Summary Request activation email
// @Description Request a new activation email to be sent to the user's email address.
// @Tags user
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string "Activation email sent successfully"
// @Failure 400 {object} map[string]string "User unauthorized"
// @Router /user/activate/resend [post]
func (controller *UserController) ResendActivation(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	if err := controller.userService.ResendActivation(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "message sent successfully",
	})
}

// DeleteUser godoc
//
// @Summary Delete user account
// @Description Permanently delete the authenticated user's account.
// @Tags user
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string "User account deleted successfully"
// @Failure 400 {object} map[string]string "Invalid user ID or deletion failed"
// @Failure 401 {object} map[string]string "User not authenticated"
// @Router /user/profile [delete]
func (controller *UserController) DeleteUser(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	if err := controller.userService.Delete(userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}

// GetUser godoc
//
// @Summary Get user data
// @Description Get the authenticated user's account data.
// @Tags user
// @Accept  json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.User
// @Failure 400 {object} map[string]string "Invalid user ID or deletion failed"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 401 {object} map[string]string "User unauthorized"
// @Router /user/profile [get]
func (controller *UserController) GetUser(ctx *gin.Context) {
	userId := ctx.GetString("user_id")
	user, err := controller.userService.GetUser(userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// ChangeEmail godoc
//
// @Summary Change user email
// @Description Reset current email, request new email confirmation
// @Tags user
// @Produce json
// @Security BearerAuth
// @Param request body	ChangeEmailRequest true "Change email request"
// @Success 200 {object} map[string]string "Email reset successfully"
// @Success 400 {object} map[string]string "Email has not changed"
// @Failure 401 {object} map[string]string "User unauthorized"
// @Router /user/email/change [post]
func (controller *UserController) ChangeEmail(ctx *gin.Context) {
	userId := ctx.GetString("user_id")

	var req *ChangeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		controller.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	if err := controller.userService.ChangeEmail(userId, req.NewEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "email changed successfully",
	})
}
