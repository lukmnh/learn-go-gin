package repository

import (
	"context"
	"errors"
	"todo_rest_api/internal/apperrors"
	"todo_rest_api/internal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(ctx context.Context, title string, completed bool) (*models.Todos, error)
	GetAll(ctx context.Context) ([]models.Todos, error)
	GetByID(ctx context.Context, id int) (*models.Todos, error)
	Update(ctx context.Context, id int, title string, completed bool) (*models.Todos, error)
	Delete(ctx context.Context, id int) error
}

type todosRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todosRepository {
	return &todosRepository{db: db}
}

func (r *todosRepository) Create(ctx context.Context, title string, completed bool) (*models.Todos, error) {
	todo := models.Todos{
		Title:     title,
		Completed: completed,
	}

	if err := r.db.WithContext(ctx).Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todosRepository) GetAll(ctx context.Context) ([]models.Todos, error) {
	todos := make([]models.Todos, 0)
	if err := r.db.WithContext(ctx).Order("created_at DESC").Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

func (r *todosRepository) GetByID(ctx context.Context, id int) (*models.Todos, error) {
	var todo models.Todos
	err := r.db.WithContext(ctx).First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return &todo, nil
}

func (r *todosRepository) Update(ctx context.Context, id int, title string, completed bool) (*models.Todos, error) {
	var todo models.Todos
	if err := r.db.WithContext(ctx).First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	todo.Title = title
	todo.Completed = completed

	if err := r.db.WithContext(ctx).Save(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func (r *todosRepository) Delete(ctx context.Context, id int) error {
	result := r.db.WithContext(ctx).Delete(&models.Todos{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}
