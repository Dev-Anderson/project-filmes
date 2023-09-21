package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Movie struct {
	MovieID int
	Title   string
	Genres  string
}

func main() {
	// Abra o arquivo CSV

	file, err := os.Open("assets/movie.csv")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Crie um leitor CSV personalizado usando vírgula como delimitador
	customReader := csv.NewReader(bufio.NewReader(file))
	customReader.Comma = ','

	// Leia e descarte a primeira linha (cabeçalho)
	_, err = customReader.Read()
	if err != nil {
		fmt.Println("Erro ao ler o cabeçalho:", err)
		return
	}

	// Use um canal para coletar os resultados das goroutines
	resultChannel := make(chan Movie)
	done := make(chan bool) // Canal de sinalização para indicar que todas as goroutines terminaram

	// Use um grupo de espera para sincronizar as goroutines
	var wg sync.WaitGroup

	// Número de goroutines a serem usadas (ajuste conforme necessário)
	numGoroutines := 20

	// Função para processar uma linha do CSV
	processLine := func(record []string) {
		defer wg.Done()
		movieID := record[0]
		title := record[1]
		genres := record[2]

		// Converte o movieID para um inteiro (você pode adicionar tratamento de erro aqui)
		movieIDInt, _ := strconv.Atoi(movieID)

		movie := Movie{
			MovieID: movieIDInt,
			Title:   title,
			Genres:  genres,
		}

		resultChannel <- movie
	}

	// Inicie as goroutines para processar as linhas do CSV
	for i := 0; i < numGoroutines; i++ {
		go func() {
			for {
				record, err := customReader.Read()
				if err != nil {
					break
				}
				wg.Add(1)
				processLine(record)
			}
			done <- true // Sinalize que esta goroutine terminou
		}()
	}

	// Inicie uma goroutine para esperar até que todas as goroutines terminem
	go func() {
		for i := 0; i < numGoroutines; i++ {
			<-done // Aguarde todas as goroutines terminarem
		}
		close(resultChannel) // Feche o canal de resultados depois que todas as goroutines terminarem
	}()

	// Imprima os resultados conforme eles chegam no canal
	for movie := range resultChannel {
		fmt.Printf("MovieID: %d, Title: %s, Genres: %s\n", movie.MovieID, movie.Title, movie.Genres)
	}
}
