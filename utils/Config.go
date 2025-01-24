package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


type Config struct {
    DB_URL string
}

func EnvBuilder() Config {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar arquivo .ENV")
    }

    return Config{
        DB_URL: os.Getenv("DATABASE_URL"),
    }
}

var Env Config = EnvBuilder();
