package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/repository"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/service"
)

type Controller struct {
	Controller *service.PautaService
}

func NewPautaController(service *service.PautaService) *Controller {
	return &Controller{
		Controller: service,
	}
}

// InitializeRoutesPauta é a função que inicializa as rotas relacionadas ao "objeto" Pauta da API
func InitializeRoutesPauta(gin *gin.Engine) {

	repo := repository.NewPautaRepository()
	service := service.NewPautaService(repo)
	controller := NewPautaController(service)

	// basePath do grupo de rotas relacionadas a Pauta
	basePath := "/api"
	v1 := gin.Group(basePath)
	{
		v1.POST("/pauta", controller.CreatePauta)
		v1.GET("/pauta/:id", controller.ListPauta)
		v1.GET("/pautas", controller.ListAllPautas)
		v1.PUT("/pauta/:id", controller.UpdatePauta)
		v1.DELETE("/pauta/:id", controller.DeletePauta)
	}
	
}