package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Movie struct {
	MovieID int
	Title   string
	Genres  string
}

func main() {
	filePath := "assets/movie.csv"
	movies, err := readMoviesFromFile(filePath)
	if err != nil {
		fmt.Println("Erro o ler o arquivo", err)
		return
	}

	printMovies(movies)
}

func readMoviesFromFile(filePath string) ([]Movie, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	customReader := csv.NewReader(bufio.NewReader(file))
	customReader.Comma = ','

	_, err = customReader.Read() // Descarta o cabeçalho
	if err != nil {
		return nil, err
	}

	var movies []Movie
	for {
		record, err := customReader.Read()
		if err != nil {
			break
		}

		movieID := record[0]
		title := record[1]
		genres := record[2]

		movieIDInt, _ := strconv.Atoi(movieID)

		movie := Movie{
			MovieID: movieIDInt,
			Title:   title,
			Genres:  genres,
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func printMovies(movies []Movie) {
	for _, movie := range movies {
		fmt.Printf("MovieID: %d, Title: %s, Genres: %s\n", movie.MovieID, movie.Title, movie.Genres)
	}
}
