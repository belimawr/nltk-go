package store

import (
	"database/sql"
	"log"

	"github.com/belimawr/nltk-go/models"
)

// Store - Interface to "Floresta Sintática" corpus
type Store interface {
	ReadArvores(limit int) []models.Arvores
	GetWord(word string) ([]models.Ramos, error)
	GetLema(lema string) ([]models.Ramos, error)
}

// NewStore - Returns a Postgres implementation of Store
func NewStore(db *sql.DB) Store {
	return &sqlStore{
		DB: db,
	}
}

// sqlStore - Postgres implementation of Store
type sqlStore struct {
	DB *sql.DB
}

func (s *sqlStore) GetWord(word string) ([]models.Ramos, error) {
	rows, err := s.DB.Query("SELECT id, arvore, palavra, lema, funcao, forma, morfo, pai FROM ramos WHERE palavra = $1", word)

	if err != nil {
		log.Printf("Error querying database: %s", err.Error())
		return []models.Ramos{}, err
	}

	defer rows.Close()

	ramos := []models.Ramos{}

	for rows.Next() {
		r := models.Ramos{}

		err := rows.Scan(
			&r.Id,
			&r.Arvore,
			&r.Palavra,
			&r.Lema,
			&r.Funcao,
			&r.Forma,
			&r.Morfo,
			&r.Pai,
		)

		if err != nil {
			log.Printf("Error reading database row: %s", err.Error())
			return ramos, err
		}

		ramos = append(ramos, r)
	}

	err = rows.Err()

	if err != nil {
		log.Printf("Error reading database rows: %q", rows.Err())
		return ramos, rows.Err()
	}

	return ramos, nil
}

func (s *sqlStore) GetLema(lema string) ([]models.Ramos, error) {
	rows, err := s.DB.Query("SELECT id, arvore, palavra, lema, funcao, forma, morfo, pai FROM ramos WHERE lema = $1", lema)

	if err != nil {
		log.Printf("Error querying database: %s", err.Error())
	}

	defer rows.Close()

	ramos := []models.Ramos{}

	for rows.Next() {
		r := models.Ramos{}

		err := rows.Scan(
			&r.Id,
			&r.Arvore,
			&r.Palavra,
			&r.Lema,
			&r.Funcao,
			&r.Forma,
			&r.Morfo,
			&r.Pai,
		)

		if err != nil {
			log.Printf("Error reading database row: %s", err.Error())
			return ramos, err
		}

		ramos = append(ramos, r)
	}

	return ramos, nil
}

func (s *sqlStore) ReadArvores(limit int) []models.Arvores {
	rows, err := s.DB.Query("SELECT id, referencia, n, cad, sec, sem, texto, analise FROM arvores LIMIT $1", limit)

	if err != nil {
		log.Printf("Error querying database: %s", err.Error())
	}

	defer rows.Close()

	arvores := []models.Arvores{}

	for rows.Next() {
		a := models.Arvores{}

		err := rows.Scan(
			&a.ID,
			&a.Referencia,
			&a.N,
			&a.Cad,
			&a.Sec,
			&a.Sem,
			&a.Texto,
			&a.Analise)

		if err != nil {
			log.Printf("Error reading database response: %s", err.Error())
			break
		}

		arvores = append(arvores, a)
	}

	return arvores
}
