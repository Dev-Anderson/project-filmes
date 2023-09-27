package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func LoadEnv() Env {
	godotenv.Load()
	return Env{
		User:     os.Getenv("db_user"),
		Password: os.Getenv("db_password"),
		Host:     os.Getenv("db_host"),
		Port:     os.Getenv("db_port"),
		Database: os.Getenv("db_database"),
	}
}
