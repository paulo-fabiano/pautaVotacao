package entity

import "strings"

type Pauta struct {

	ID int
	Nome string
	Descricao string
	
}

func (pauta Pauta) Validar() {

	pauta.formatarEntidade(&pauta)

}

func (p Pauta) formatarEntidade(pauta *Pauta) {

	pauta.Nome = strings.TrimSpace(pauta.Nome)
	pauta.Descricao = strings.TrimSpace(pauta.Descricao)

}