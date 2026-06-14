package service

import (
	"context"
	"strings"
	"time"
	"todo_rest_api/internal/dto"
	"todo_rest_api/internal/models"
	"todo_rest_api/internal/repository"
)

type TodoService interface {
	CreateTodo(ctx context.Context, req dto.CreateTodoReq) (*dto.TodoResponse, error)
	GetAllTodos(ctx context.Context) ([]dto.TodoResponse, error)
	GetTodoByID(ctx context.Context, id int) (*dto.TodoResponse, error)
	UpdateTodo(ctx context.Context, id int, req dto.UpdateTodoReq) (*dto.TodoResponse, error)
	DeleteTodo(ctx context.Context, id int) error
}

type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *todoService {
	return &todoService{repo: repo}
}

const dbTimeout = time.Second * 5

func (s *todoService) CreateTodo(ctx context.Context, req dto.CreateTodoReq) (*dto.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	title := strings.TrimSpace(req.Title)

	todo, err := s.repo.Create(ctx, title, req.Completed)
	if err != nil {
		return nil, err
	}
	return toResponse(todo), nil
}

func (s *todoService) GetAllTodos(ctx context.Context) ([]dto.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	todos, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]dto.TodoResponse, 0, len(todos))
	for _, t := range todos {
		res = append(res, *toResponse(&t))
	}
	return res, nil
}

func (s *todoService) GetTodoByID(ctx context.Context, id int) (*dto.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return toResponse(todo), nil
}

func (s *todoService) UpdateTodo(ctx context.Context, id int, req dto.UpdateTodoReq) (*dto.TodoResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	title := existing.Title
	if req.Title != nil {
		title = strings.TrimSpace(*req.Title)
	}

	completed := existing.Completed
	if req.Completed != nil {
		completed = *req.Completed
	}

	todo, err := s.repo.Update(ctx, id, title, completed)
	if err != nil {
		return nil, err
	}
	return toResponse(todo), nil
}

func (s *todoService) DeleteTodo(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	defer cancel()

	return s.repo.Delete(ctx, id)
}

func toResponse(t *models.Todos) *dto.TodoResponse {
	return &dto.TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}
