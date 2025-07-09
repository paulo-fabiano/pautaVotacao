package config

import (

	"github.com/paulo-fabiano/pautaVotacao/internal/config/database"
	"github.com/paulo-fabiano/pautaVotacao/internal/config/server"
)

func SetupConfigAPI() error {

	err := database.ConnectDatabase()
	if err != nil {
		panic("Erro ao inicializar o banco: " + err.Error())
	}

	err = server.InitializeServer()
	if err != nil {
		panic("Erro ao incializar o server: " + err.Error())
	}

	return nil

}