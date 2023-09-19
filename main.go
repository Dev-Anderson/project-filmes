package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("assets/movie.csv")
	if err != nil {
		log.Panic("Falha ao abrir o arquivo", err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		moveID := record[0]
		title := record[1]
		genres := record[2]

		fmt.Printf("moveID: %s, Title: %s, Genres: %s\n", moveID, title, genres)
	}

}
