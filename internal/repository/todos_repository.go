package repository

import (
	"context"
	"fmt"
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

func GetAllTodos(pool *pgxpool.Pool) ([]models.Todos, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `SELECT id, title, completed, created_at, updated_at 
              FROM golang.todos_user 
              ORDER BY created_at DESC`
	rows, err := pool.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := make([]models.Todos, 0)
	for rows.Next() {
		var todo models.Todos
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodoById(pool *pgxpool.Pool, id int) (*models.Todos, error) {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `SELECT id, title, completed, created_at, updated_at
				FROM golang.todos_user
				WHERE id = $1`

	var todos models.Todos
	err := pool.QueryRow(ctx, query, id).Scan(
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

func UpdateTodo(pool *pgxpool.Pool, id int, title string, completed bool) (*models.Todos, error) {
	var ctx context.Context
	var cancel context.CancelFunc

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `UPDATE golang.todos_user
				SET title = $1, completed = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3
				RETURNING id, title, completed, created_at, updated_at`

	var todos models.Todos
	err := pool.QueryRow(ctx, query, title, completed, id).Scan(
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

func DeleteTodo(pool *pgxpool.Pool, id int) error {
	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	query := `DELETE FROM golang.todos_user
				where id = $1`

	var commandTag, err = pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("todo with id %d does not exist", id)
	}

	return nil
}
