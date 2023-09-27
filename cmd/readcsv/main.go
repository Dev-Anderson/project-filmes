package main

import (
	"filmes/internal/database"
	"fmt"
)

func main() {
	fmt.Println("Conectando com o banco dados")
	database.ConnectDatabase()
}
