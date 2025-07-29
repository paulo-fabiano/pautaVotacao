package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/paulo-fabiano/pautaVotacao/internal/app/pauta/entity"
	"github.com/paulo-fabiano/pautaVotacao/internal/config/database"
)

type Repository struct {
	db *sql.DB
}

func NewPautaRepository() *Repository {
	return &Repository{db: database.GetConnectionDatabase()}
}

// Create é a função do Repository que cria uma pauta
func(r Repository) Create(pauta entity.Pauta) (uint64, error) {

	query := "INSERT INTO t_pauta_votacao (nome, descricao) VALUES ($1, $2) RETURNING id;"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	defer stmt.Close()

	var ID int
	err = stmt.QueryRow(&pauta.Nome, &pauta.Descricao).Scan(&ID)
	if err != nil {
		log.Println(err)
		return 0, errors.New("erro ao salvar objeto no banco de dados")
	}
	
	return uint64(ID), nil
 
}

// Get é a função do Repository que busca uma pauta
func(r Repository) Get(id uint64) (entity.Pauta, error) {
	
	query := "SELECT * FROM t_pauta_votacao WHERE id = $1;"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return entity.Pauta{}, err
	}
	defer stmt.Close()

	var pauta entity.Pauta
	err = stmt.QueryRow(id).Scan(&pauta.ID, &pauta.Nome, &pauta.Descricao)
	if err != nil {
		log.Println("Erro ao executar a busca no banco de dados", err)
		return pauta, err
	}

	return pauta, nil

}

// GetAll é a função do Repository busca todas as pautas
func(r Repository) GetAll() ([]entity.Pauta, error) {
	
	query := "SELECT * FROM t_pauta_votacao;"
	
	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPautas []entity.Pauta
	for rows.Next() {

		var pauta entity.Pauta
		err := rows.Scan(&pauta.ID, &pauta.Nome, &pauta.Descricao)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		listPautas = append(listPautas, pauta)

	}

	return listPautas, nil

}

// Update é a função do Repository que atualiza uma Pauta
func(r Repository) Update(id uint64, data interface{}) error {

	queryExist := "SELECT EXISTS(SELECT id FROM tasks WHERE id = $1)"

	stmtExist, err := r.db.Prepare(queryExist)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmtExist.Close()

	var existe bool
	err = stmtExist.QueryRow(id).Scan(&existe)
	if err != nil {
		log.Println(err)
		return err
	}

	if !existe {
		return fmt.Errorf("ID não existe no banco de dados")
	}

	var pauta entity.Pauta
	pauta, err = r.Get(id)
	if err != nil {
		log.Println(err)
		return err
	}

	queryUpdate := "UPDATE t_pauta SET nome = $1, descricao = $2 WHERE id = $3;"
	_, err = r.db.Exec(queryUpdate, pauta.Nome, pauta.Descricao, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
	
}

// Delete é a função do Repository que apaga uma Pauta
func(r Repository) Delete(id uint64) error {
	
	queryExist := "SELECT EXISTS(SELECT id FROM tasks WHERE id = $1)"

	stmt, err := r.db.Prepare(queryExist)
	if err != nil {
		log.Println(err)
		return err
	}
	defer stmt.Close()

	var existe bool
	err = stmt.QueryRow(id).Scan(&existe)
	if err != nil {
		log.Println(err)
		return err
	}

	if !existe {
		return fmt.Errorf("ID não existe no banco de dados")
	}

	queryDelete := "DELETE FROM t_pautas WHERE id = $1;"

	stmt, err = r.db.Prepare(queryDelete)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil

}