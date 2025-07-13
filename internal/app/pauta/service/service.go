package service

import (
	"fmt"
	"log"

	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/dto"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/entity"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/repository"
)

type PautaService struct {

	repo *repository.Repository

}

func NewPautaService(r *repository.Repository) *PautaService {

	return &PautaService{repo: r}

}

func validadePauta(pauta entity.Pauta) error {

	if pauta.Nome == "" {
		return fmt.Errorf("O campo Nome está vazio")
	}

	if pauta.Descricao == "" {
		return fmt.Errorf("O campo Descricao está vazio")
	}

	return nil

}

func transformaEmEntity(dto dto.PautaRequest) entity.Pauta {

	var pauta entity.Pauta
	pauta.Nome = dto.Nome
	pauta.Descricao = dto.Descricao
	return pauta

}

func (r *PautaService) Create(dto dto.PautaRequest) (entity.Pauta, error) {

	pauta := transformaEmEntity(dto)

	err := validadePauta(pauta) 
	if err != nil {
		return entity.Pauta{}, err
	}

	id, err := r.repo.Create(pauta)
	if err != nil {
		log.Println(err)
		return entity.Pauta{}, err
	}

	pautaCreated, err := r.repo.Get(id)

	return pautaCreated, nil

}

func (r *PautaService) List(id int) (entity.Pauta, error) {

	// if id ==  {
	// 	return entity.Pauta{}, fmt.Errorf("ID é null")
	// }

	var pauta entity.Pauta
	pauta, err := r.repo.Get(id)
	if err != nil {
		return entity.Pauta{}, err
	}

	return pauta, nil

}

func (r *PautaService) ListAll() ([]entity.Pauta, error) {

	listaPautas, err := r.repo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return listaPautas, nil

}

func (r *PautaService) Update(id int, pauta entity.Pauta) error {

	// Implements this logic after

	return nil

}

func (r *PautaService) Delete(id int) error {

	err := r.repo.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}