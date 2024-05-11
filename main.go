package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Ponto de entrada
func main() {

	// Carrega as variáveis de ambiente e verifica se houver alguma falha
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("TEST"))
}
