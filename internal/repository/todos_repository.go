package repository

import (
	"context"
	"errors"
	"todo_rest_api/internal/apperrors"
	"todo_rest_api/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TodoRepository interface {
	Create(ctx context.Context, title string, completed bool) (*models.Todos, error)
	GetAll(ctx context.Context) ([]models.Todos, error)
	GetByID(ctx context.Context, id int) (*models.Todos, error)
	Update(ctx context.Context, id int, title string, completed bool) (*models.Todos, error)
	Delete(ctx context.Context, id int) error
}

type todosRepository struct {
	pool *pgxpool.Pool
}

func NewTodoRepository(pool *pgxpool.Pool) *todosRepository {
	return &todosRepository{pool: pool}
}

func (r *todosRepository) Create(ctx context.Context, title string, completed bool) (*models.Todos, error) {
	query := `INSERT INTO golang.todos_user (title, completed) VALUES ($1, $2)
			  RETURNING id, title, completed, created_at, updated_at`
	var todo models.Todos
	err := r.pool.QueryRow(ctx, query, title, completed).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todosRepository) GetAll(ctx context.Context) ([]models.Todos, error) {
	query := `SELECT id, title, completed, created_at, updated_at FROM golang.todos_user ORDER BY created_at DESC`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]models.Todos, 0)
	for rows.Next() {
		var todo models.Todos
		if err := rows.Scan(
			&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
		); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todosRepository) GetByID(ctx context.Context, id int) (*models.Todos, error) {
	query := `SELECT id, title, completed, created_at, updated_at
				FROM golang.todos_user
				WHERE id = $1`

	var todo models.Todos
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todosRepository) Update(ctx context.Context, id int, title string, completed bool) (*models.Todos, error) {
	query := `UPDATE golang.todos_user
				SET title = $1, completed = $2, updated_at = CURRENT_TIMESTAMP
				WHERE id = $3
				RETURNING id, title, completed, created_at, updated_at`

	var todo models.Todos
	err := r.pool.QueryRow(ctx, query, title, completed, id).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todosRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM golang.todos_user WHERE id = $1`

	commandTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}
