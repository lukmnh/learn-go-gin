package repository

import (
	"context"
	"time"
	"todo_rest_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodos(pool *pgxpool.Pool, title string, completed bool) (*models.Todos, error) {
	// context akan selalu digunakan untuk I/O db atau http
	var ctx context.Context
	var cancel context.CancelFunc
	// buat timeout hit db 5 detik
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `INSERT INTO golang.todos_user (title, completed)
						VALUES ($1, $2)
						RETURNING id, title, completed, created_at, updated_at`

	var todos models.Todos
	err := pool.QueryRow(ctx, query, title, completed).Scan(
		&todos.ID,
		&todos.Title,
		&todos.Completed,
		&todos.CreatedAt,
		&todos.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &todos, nil
}
