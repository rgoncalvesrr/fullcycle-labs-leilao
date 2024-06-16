package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Não foi possível ler variáveis de ambiente. Abortando aplicação")
		return
	}
}
