package router

import (

	"fmt"
	"github.com/gin-gonic/gin"
	pautaController "github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/controller"

)

func InitializeRoutesApp(ginRouter *gin.Engine) error {

	if ginRouter == nil {
		return fmt.Errorf("ginRouter is nil")
	}

	pautaController.InitializeRoutesPauta(ginRouter)

	return nil

}