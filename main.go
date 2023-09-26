package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

type Movie struct {
	MovieID string
	Title   string
	Genres  string
}

func main() {
	var csvFile = strings.NewReader(`"movieId","title","genres"`)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ','

	var filmes []Movie

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		filmes = append(filmes, Movie{
			MovieID: line[0],
			Title:   line[1],
			Genres:  line[2],
		})
	}

	fmt.Println(filmes)

}
