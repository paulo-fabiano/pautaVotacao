package router

import (

	"fmt"
	"github.com/gin-gonic/gin"
	pautaController "github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/controller"

)

// InitializeRoutesApp recebe um *gin.Engine e inicializa as rotas da API
func InitializeRoutesApp(ginRouter *gin.Engine) error {

	if ginRouter == nil {
		return fmt.Errorf("ginRouter is nil")
	}

	pautaController.InitializeRoutesPauta(ginRouter)

	return nil

}