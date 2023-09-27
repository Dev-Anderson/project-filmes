package database

import (
	"database/sql"
	"filmes/cmd/config"
	"fmt"
)

func ConnectDatabase() (*sql.DB, error) {
	stringConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		config.LoadEnv().Host, config.LoadEnv().User, config.LoadEnv().Password, config.LoadEnv().Database)

	db, err := sql.Open("postgres", stringConnection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Falha ao realizar o ping no banco de dados", err)
	}

	return db, nil

}
