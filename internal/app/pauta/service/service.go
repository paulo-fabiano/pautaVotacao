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
		return fmt.Errorf("o campo Nome está vazio")
	}

	if pauta.Descricao == "" {
		return fmt.Errorf("o campo Descricao está vazio")
	}

	return nil

}

func transformaEmEntity(dto dto.PautaRequest) entity.Pauta {

	var pauta entity.Pauta
	pauta.Nome = dto.Nome
	pauta.Descricao = dto.Descricao
	return pauta

}

// Create é a função do Service para criar uma Pauta
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

// List é a função do Service para listar uma Pauta
func (r *PautaService) List(id uint64) (entity.Pauta, error) {

	var pauta entity.Pauta
	pauta, err := r.repo.Get(id)
	if err != nil {
		return entity.Pauta{}, err
	}

	return pauta, nil

}

// ListAll é a função do Service para listar pautas
func (r *PautaService) ListAll() ([]entity.Pauta, error) {

	listaPautas, err := r.repo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return listaPautas, nil

}

// Update é a função do Service para atualizar uma Pauta
func (r *PautaService) Update(id uint64, pauta entity.Pauta) error {

	err := r.repo.Update(id, pauta)
	if err != nil {
		return err
	}

	return nil

}

// Delete é a função do Service para apagar uma Pauta
func (r *PautaService) Delete(id uint64) error {

	err := r.repo.Delete(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}