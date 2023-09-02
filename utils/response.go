package util

import (
	"github.com/gin-gonic/gin"
)

type Responses struct {
	Success    bool         `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorResponse struct {
	Success    bool         `json:"success"`
	Method     string      `json:"method"`
	Error      interface{} `json:"error"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := Responses{
		Success:    true,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := ErrorResponse{
		Success:    false,
		Method:     Method,
		Error:      Error,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}
