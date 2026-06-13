package main

import (
	"log"
	"todo_rest_api/internal/config"
	"todo_rest_api/internal/database"
	"todo_rest_api/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	var pool *pgxpool.Pool
	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	defer pool.Close() // jalanin ini ketika function selesai jalan

	/*
		gin.Default() membuat object router untuk menerima HTTP Request.
	*/
	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "TODOS RestAPI",
			"status":   "200",
			"database": "connected",
		})
	})
	router.POST("/todos", handlers.CreateTodoHandler(pool))
	router.GET("/todos", handlers.GetAllTodoHandler(pool))
	router.GET("/:id", handlers.GetTodoByIdHandler(pool))
	router.PUT("/:id", handlers.UpdateTodoByIdHandler(pool))
	router.DELETE("/:id", handlers.DeleteTodoByIdHandler(pool))

	router.Run(":" + cfg.PORT)
}
