package store

import (
	"database/sql"
	"log"

	"github.com/belimawr/nltk-go/models"
)

type Store struct {
	DB *sql.DB
}

func (s *Store) ReadArvores(limit int) []models.Arvores {
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
