package handlers

import (
	"net/http"
	"todo_rest_api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateTodoReq struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

func CreateTodoHandler(pool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateTodoReq

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		todo, err := repository.CreateTodos(pool, req.Title, req.Completed)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusCreated, todo)
	}
}
