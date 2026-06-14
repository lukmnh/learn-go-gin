package main

import (
	"log"
	"todo_rest_api/internal/bootstrap"
	"todo_rest_api/internal/config"
	"todo_rest_api/internal/database"
	"todo_rest_api/internal/router"
	"todo_rest_api/internal/util/jwt"
	"todo_rest_api/internal/util/validation"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	jwt.Init(cfg.JWTSecret)

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer func() {
		sqlDB, err := pool.DB()
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
		sqlDB.Close()
	}()

	if err := validation.RegisterValidation(); err != nil {
		log.Fatal("Failed to register validation", err)
	}

	handlers := bootstrap.BuildHandler(pool)
	r := router.Setup(handlers)

	if err := r.Run(":" + cfg.PORT); err != nil {
		log.Fatal("Failed to start server", err)
	}
}
