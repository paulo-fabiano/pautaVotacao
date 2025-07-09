package service

import (

	"fmt"
	"log"

	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/entity"
	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/repository"

)

type PautaService struct {

	repo *repository.Repository

}

func NewPautaService(r *repository.Repository) *PautaService {

	return &PautaService{repo: r}

}

func validadePauta(pauta *entity.Pauta) error {

	if pauta.Nome == "" {
		return fmt.Errorf("O campo Nome está vazio")
	}

	if pauta.Descricao == "" {
		return fmt.Errorf("O campo Descricao está vazio")
	}

	return nil

}

func (r *PautaService) Create(pauta *entity.Pauta) (*entity.Pauta, error) {

	if pauta == nil {
		return nil, fmt.Errorf("O objeto está vazio")
	}

	err := validadePauta(pauta) 
	if err != nil {
		return nil, err
	}

	id, err := r.repo.Create(pauta)
	if err != nil {
		return nil, err
	}

	var pautaCreated *entity.Pauta
	pautaCreated, err = r.repo.Get(id)

	return pautaCreated, nil

}

func (r *PautaService) List(id *int) (*entity.Pauta, error) {

	if id == nil {
		return nil, fmt.Errorf("ID é null")
	}

	var pauta *entity.Pauta
	pauta, err := r.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return pauta, nil

}

func (r *PautaService) ListAll() (*[]entity.Pauta, error) {

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

func (r *PautaService) Delete(id *int) error {

	if id == nil {
		return fmt.Errorf("ID is null")
	}

	err := r.repo.Delete(*id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}