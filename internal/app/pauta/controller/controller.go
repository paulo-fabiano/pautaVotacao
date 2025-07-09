package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/entity"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/service"
	"github.com/paulo-fabiano/pautaVotacao/internal/utils/handler"
)

type Controller struct {
	Controller *service.PautaService
}

func NewPautaController(service *service.PautaService) *Controller {
	return &Controller{
		Controller: service,
	}
}

func (c *Controller) CreatePauta(ctx *gin.Context) {

	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "msg",
		"errorCode": "code",
	})	
	
	pauta := entity.Pauta{ID: 2, Nome: "Teste", Descricao: "Teste"}
	c.Controller.Create(&pauta)
	// if request.Method != "POST" {
	// 	handler.SendError(writer, http.StatusBadRequest, "método não permitido")
	// 	return
	// }

	// repository := repository.NewPautaRepository()
	// service := service.NewPautaService(repository)

	// var pauta entity.Pauta
	// pautaDecoder := json.NewDecoder(request.Body)
	// err := pautaDecoder.Decode(&pauta)
	// if err != nil {
	// 	handler.SendError(writer, http.StatusInternalServerError, "erro interno")
	// 	return
	// }

	// // Implements return obj created
	// _, err = service.Create(&pauta)
	// if err != nil {
	// 	handler.SendError(writer, http.StatusInternalServerError, "erro interno")
	// 	return
	// }

	// handler.SendSucess(writer, http.StatusOK, "created")

}

func ListPauta(ctx *gin.Context) {

	// if request.Method != "GET" {
	// 	handler.SendError(writer, http.StatusBadRequest, "método não permitido")
	// 	return
	// }

	// repository := repository.NewPautaRepository()
	// service := service.NewPautaService(repository)

	// paramID := request.URL.Query().Get("id")
	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	handler.SendError(writer, http.StatusInternalServerError, "erro interno")
	// 	return
	// }

	// var pauta *entity.Pauta
	// pauta, err = service.List(&id)
	// if err != nil {
	// 	handler.SendError(writer, http.StatusInternalServerError, "erro interno")
	// 	return
	// }

	// handler.SendSucess(writer, http.StatusOK, pauta)
	
}


func (c Controller) ListAllPautas(ctx *gin.Context) {

	listaPautas, err := c.Controller.ListAll()
	if err != nil {
		log.Println(err)
		handler.SendError(ctx, http.StatusInternalServerError, "Erro interno")
		return
	}

	type Data struct {
		Data *[]entity.Pauta
	}
	handler.SendSucess(ctx, http.StatusOK, Data{Data: listaPautas})
	
}

func UpdatePauta(ctx *gin.Context) {

	// if request.Method != "PUT" || request.Method != "PATCH" {
	// 	handler.SendError(writer, http.StatusBadRequest, "método não permitido")
	// 	return
	// }

	// // implementar depois
	
}

func DeletePauta(ctx *gin.Context) {

	// if request.Method != "GET" {
	// 	handler.SendError(writer, http.StatusBadRequest, "método não permitido")
	// 	return
	// }
	
	// repository := repository.NewPautaRepository()
	// service := service.NewPautaService(repository)

	// paramID := request.URL.Query().Get("id")
	// id, err := strconv.Atoi(paramID)
	// if err != nil {
	// 	handler.SendError(writer, http.StatusInternalServerError, "erro interno")
	// 	return 
	// }

	// err = service.Delete(&id)
	// if err != nil {
	// 	handler.SendError(writer, http.StatusInternalServerError, "erro interno")
	// 	return
	// }

	// handler.SendSucess(writer, http.StatusNoContent, "deleted with sucess")
	
}