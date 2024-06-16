package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/rgoncalvesrr/fullcycle-labs-leilao/config/database/mongodb"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Não foi possível ler variáveis de ambiente. Abortando aplicação")
		return
	}

	_, err := mongodb.NewMongoDBConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

}
