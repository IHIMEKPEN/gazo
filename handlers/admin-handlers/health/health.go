package handlerHealth

import (
	"net/http"

	util "gazo/utils"

	"github.com/gin-gonic/gin"
)

func HealthHandler(ctx *gin.Context) {
	util.APIResponse(ctx, util.GodotEnv("APP_NAME")+" API Running...", http.StatusCreated, http.MethodPost, nil)

}
