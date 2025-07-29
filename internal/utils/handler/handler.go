package handler

import (

	"time"
	"github.com/gin-gonic/gin"

)

// ErrorResponse é a struct para mensagens de erro da API
type ErrorResponse struct {

	Message string `json:"message"`
	DateTime time.Time `json:"date"`

}

// SendError é a função que retorna erros da API
func SendError(ctx *gin.Context, code int, errorMessage string) {

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, ErrorResponse{
		Message:  errorMessage,
		DateTime: time.Now(),
	})

}

// SendSucess é a função que retorna mensagens de sucesso da API
func SendSucess(ctx *gin.Context, code int, data interface{}) {

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, data)

}