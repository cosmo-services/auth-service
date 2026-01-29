package ping_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	msg string
}

func NewPingController() PingController {
	return PingController{
		msg: "Pong!",
	}
}

// Ping godoc
//
//	@Summary		Ping endpoint
//	@Description	Check if the service is alive
//	@Tags			health
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Successful response"
//	@Router			/api/v1/ping [get]
func (c PingController) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": c.msg,
	})
}
