package main

import (
	"log"
	"todo_rest_api/internal/bootstrap"
	"todo_rest_api/internal/config"
	"todo_rest_api/internal/database"
	"todo_rest_api/internal/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer pool.Close()

	handlers := bootstrap.BuildHandler(pool)
	r := router.Setup(handlers)

	if err := r.Run(":" + cfg.PORT); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
