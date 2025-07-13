package dto

type PautaRequest struct {

	Nome string	`json:"nome"`
	Descricao string `json:"descricao"`
	
}

type PautaResponse struct {

	ID int	`json:"id"`
	Nome string	`json:"nome"`
	Descricao string `json:"descricao"`
	
}