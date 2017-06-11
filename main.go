package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/belimawr/nltk-go/config"
	"github.com/belimawr/nltk-go/store"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", config.PostgresConnectionString())

	if err != nil {
		log.Printf("Error connecting to database: %s", err.Error())
		os.Exit(1)
	}

	store := store.Store{DB: db}

	for _, a := range store.ReadArvores(10) {
		fmt.Printf("%#v\n", a)
	}
}
