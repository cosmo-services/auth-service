package password_api

import (
	"main/internal/domain/password"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PasswordController struct {
	pswdService *password.PasswordService
}

func NewPasswordController(pswdService *password.PasswordService) *PasswordController {
	return &PasswordController{
		pswdService: pswdService,
	}
}

// ValidatePassword godoc
//
// @Summary Validate new password
// @Description Validate password string for special characters
// @Tags password
// @Produce json
// @Param password query string true "Passwod string"
// @Success 200 {object} map[string]string "Password is ok"
// @Failure 400 {object} map[string]string "Password param is required"
// @Failure 422 {object} map[string]string "Password incorrect"
// @Router /password/validate [get]
func (controller *PasswordController) ValidatePassword(ctx *gin.Context) {
	password, ok := ctx.GetQuery("password")
	if !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password param is required"})
		return
	}

	if err := controller.pswdService.ValidatePassword(password); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
