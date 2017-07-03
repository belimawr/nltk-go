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

	floresta := store.NewStore(db)

	fmt.Println(`Reading some "Árvores":`)

	arvores, err := floresta.ReadArvores(10)
	if err != nil {
		panic(err.Error())
	}

	for _, a := range arvores {
		fmt.Printf("%#v\n", a)
	}

	fmt.Println("\n\n", `Reading some "Ramos" (palavra="anda"):`)

	ramos, err := floresta.GetWord("anda")
	if err != nil {
		panic(err.Error())
	}

	for _, r := range ramos {
		fmt.Printf("%#v\n", r)
	}

	fmt.Println("\n\n", `Reading some "Ramos" (lema="andar"):`)

	lemas, err := floresta.GetLema("andar")
	if err != nil {
		panic(err.Error())
	}

	for _, r := range lemas {
		fmt.Printf("%#v\n", r)
	}
}
