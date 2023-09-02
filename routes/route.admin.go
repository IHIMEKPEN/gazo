package route

import (
	"github.com/gin-gonic/gin"
	handlerHealth "gazo/handlers/admin-handlers/health"


)

func AdminRoutes(route *gin.Engine) {

	/**
	@description All Handler Admin
	*/
	// adminHandler := handlerHealth.HealthHandler(ctx *gin.Context)

	
	/**
	@description All Admin Route
	*/
	groupRoute := route.Group("")
	groupRoute.GET("/", handlerHealth.HealthHandler)

}
