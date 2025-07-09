package repository

import (
	"database/sql"
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

func(r Repository) Create(pauta *entity.Pauta) (*int, error) {

	query := "INSERT INTO %s (nome, descricao) VALUES ($1, $2) RETURNING id;"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var id int
	err = stmt.QueryRow(&pauta.Nome, &pauta.Descricao).Scan(&id)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("Erro ao salvar objeto no banco de dados")
	}
	
	return &id, nil
 
}

func(r Repository) Get(id *int) (*entity.Pauta, error) {
	
	query := "SELECT * FROM t_pauta WHERE id = $1;"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer stmt.Close()

	var pauta entity.Pauta
	err = stmt.QueryRow(id).Scan(&pauta.ID, &pauta.Nome, &pauta.Descricao)
	if err != nil {
		log.Println("Erro ao executar a busca no banco de dados")
		return nil, err
	}

	return &pauta, nil

}

func(r Repository) GetAll() (*[]entity.Pauta, error) {
	
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

	return &listPautas, nil

}

func(r Repository) Update(id int, data interface{}) error {

	queryExist := fmt.Sprintf("SELECT EXISTS(SELECT id FROM tasks WHERE id = $1)")

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

	var pauta *entity.Pauta
	pauta, err = r.Get(&id)
	if err != nil {
		log.Println(err)
		return err
	}

	queryUpdate := fmt.Sprintf("UPDATE t_pauta SET nome = $1, descricao = $2 WHERE id = $3;")
	_, err = r.db.Exec(queryUpdate, pauta.Nome, pauta.Descricao, id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
	
}

func(r Repository) Delete(id int) error {
	
	queryExist := fmt.Sprintf("SELECT EXISTS(SELECT id FROM tasks WHERE id = $1)")

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

	queryDelete := fmt.Sprintf("DELETE FROM t_pautas WHERE id = $1;")

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