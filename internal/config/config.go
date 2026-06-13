package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	PORT        string
}

func Load() (*Config, error) {
	var err error = godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	/*
		Function ini menerima object Config asli (by reference), bukan copy

		&config → ambil alamatnya
		* → buka isi dari alamat tersebut
	*/
	var config *Config = &Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		PORT:        os.Getenv("PORT"),
	}
	return config, nil
}
