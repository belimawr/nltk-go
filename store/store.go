package store

import (
	"database/sql"
	"log"

	"github.com/belimawr/nltk-go/models"
)

// NewSQLStore - Returns an implementation of Store that uses a SQL database
func NewSQLStore(db *sql.DB) Store {
	return sqlStore{DB: db}
}

// Store - Interface to read "Floresta Sintática"
type Store interface {
	ReadArvores(limit int) []models.Arvores
	GetRamo(word string) models.Ramos
}

type sqlStore struct {
	DB *sql.DB
}

func (s sqlStore) GetRamo(word string) models.Ramos {
	query := "SELECT id, arvore, palavra, lema, funcao, forma, morfo, pai FROM ramos WHERE LOWER(palavra) = LOWRE(?);"

	row := s.DB.QueryRow(query, word)

	ramo := models.Ramos{}

	err := row.Scan(
		&ramo.Id,
		&ramo.Arvore,
		&ramo.Palavra,
		&ramo.Lema,
		&ramo.Funcao,
		&ramo.Forma,
		&ramo.Morfo,
		&ramo.Pai)

	if err != nil {
		log.Printf("Error querying database: %s", err.Error())
	}

	return ramo
}

func (s sqlStore) ReadArvores(limit int) []models.Arvores {
	rows, err := s.DB.Query("SELECT id, referencia, n, cad, sec, sem, texto, analise FROM arvores LIMIT $1", limit)

	if err != nil {
		log.Printf("Error querying database: %s", err.Error())
	}

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
