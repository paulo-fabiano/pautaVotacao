package controller

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/dto"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/entity"
	"github.com/paulo-fabiano/pautaVotacao/internal/utils/handler"
)

// CreatePauta é a função que cria uma nova Pauta no banco de dados
func (c *Controller) CreatePauta(ctx *gin.Context) {
	
	var data dto.PautaRequest
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		handler.SendError(ctx, http.StatusInternalServerError, "erro interno do servidor")
		return
	}

	pautaCreated, err := c.Controller.Create(data)
	if err != nil {
		log.Println(err)
		handler.SendError(ctx, http.StatusInternalServerError, "Erro interno do servidor")
		return
	}

	type DataRes struct {
		Data entity.Pauta `json:"data"`
		Time time.Time	`json:"datetime"`
	}

	handler.SendSucess(ctx, http.StatusOK, DataRes{
		Data: pautaCreated,
		Time: time.Now(),
	})

}

func (c *Controller) ListPauta(ctx *gin.Context) {

	idString := ctx.Param("id")
	ID, err := strconv.ParseUint(idString, 10, 64)
	log.Println(idString, ID)
	if err != nil {
		handler.SendError(ctx, http.StatusBadGateway, "Erro interno do servidor")
		log.Println(err)
		return
	}

	pauta, err := c.Controller.List(ID)
	if err != nil {
		handler.SendError(ctx, http.StatusBadGateway, "Erro interno do servidor")
		log.Println(err)
		return
	}

	handler.SendSucess(ctx, http.StatusOK, pauta)
	
}


func (c *Controller) ListAllPautas(ctx *gin.Context) {

	listaPautas, err := c.Controller.ListAll()
	if err != nil {
		log.Println(err)
		handler.SendError(ctx, http.StatusInternalServerError, "Erro interno")
		return
	}

	type Data struct {
		Data []entity.Pauta `json:"data"`
	}
	handler.SendSucess(ctx, http.StatusOK, Data{Data: listaPautas})
	
}

func (c *Controller) UpdatePauta(ctx *gin.Context) {

	idString := ctx.Param("id")
	if idString == "" {
		handler.SendError(ctx, http.StatusBadRequest, "o ID está vazio")
		return
	}

	ID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "erro interno da aplicação")
	}

	var pauta entity.Pauta
	err = ctx.ShouldBindJSON(&pauta)
	if err != nil {
		handler.SendError(ctx, http.StatusBadRequest, "erro no corpo da requisição")
		return
	}

	err = c.Controller.Update(ID, pauta)
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "erro interno do servidor")
		return
	}

	handler.SendSucess(ctx, 204, nil)
	
}

func (c *Controller) DeletePauta(ctx *gin.Context) {

	idString := ctx.Param("id")
	if idString == "" {
		handler.SendError(ctx, http.StatusBadGateway, "Erro ID é null")
		return
	}

	ID, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		handler.SendError(ctx, http.StatusInternalServerError, "Erro interno do servidor")
		return
	}

	err = c.Controller.Delete(ID)
	if err != nil {
		handler.SendError(ctx, http.StatusBadGateway, "Erro interno do servidor")
		return
	}

	handler.SendSucess(ctx, http.StatusNoContent, "deleted with sucess")
	
}