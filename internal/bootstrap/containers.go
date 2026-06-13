package bootstrap

import (
	"todo_rest_api/internal/handlers"
	"todo_rest_api/internal/repository"
	"todo_rest_api/internal/router"
	"todo_rest_api/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func BuildHandler(pool *pgxpool.Pool) *router.TodoHandler {
	todoRepo := repository.NewTodoRepository(pool)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	return &router.TodoHandler{
		Todo: todoHandler,
	}
}
