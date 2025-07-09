package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/paulo-fabiano/pautaVotacao/internal/config/router"
)

var (

	GinRouter *gin.Engine

)

func InitializeServer() error {

	var (

		serverPort = os.Getenv("SERVER_PORT")

	)

	if serverPort == "" {
		serverPort = ":8080"
	}

	GinRouter = gin.Default()
	err := router.InitializeRoutesApp(GinRouter)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	GinRouter.Run(":"+serverPort)

	return nil

}