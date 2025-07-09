package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/repository"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/service"
)

func InitializeRoutesPauta(gin *gin.Engine) {

	repo := repository.NewPautaRepository()
	service := service.NewPautaService(repo)
	controller := NewPautaController(service)

	basePath := "/api"
	v1 := gin.Group(basePath)
	{
		v1.POST("/pauta", controller.CreatePauta)
		v1.GET("/pauta/:id", ListPauta)
		v1.GET("/pautas", controller.ListAllPautas)
		v1.PUT("/pauta/:id", UpdatePauta)
		v1.DELETE("/pauta/:id", DeletePauta)
	}
	
}