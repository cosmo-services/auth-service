package auth

import (
	"errors"
	"main/internal/domain/auth"

	"github.com/gin-gonic/gin"
)

func GetUserFromContext(ctx *gin.Context) (*auth.JwtPayload, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		return nil, errors.New("user not found in context")
	}

	user, ok := userValue.(*auth.JwtPayload)
	if !ok {
		return nil, errors.New("invalid user type in context")
	}

	return user, nil
}
