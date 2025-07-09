package handler

import (

	"time"
	"github.com/gin-gonic/gin"

)

type ErrorResponse struct {

	Message string `json:"message"`
	DateTime time.Time `json:"date"`

}

func SendError(ctx *gin.Context, code int, errorMessage string) {

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, ErrorResponse{
		Message:  errorMessage,
		DateTime: time.Now(),
	})

}

func SendSucess(ctx *gin.Context, code int, data interface{}) {

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, data)

}