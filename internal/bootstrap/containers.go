package bootstrap

import (
	"todo_rest_api/internal/handlers"
	"todo_rest_api/internal/repository"
	"todo_rest_api/internal/router"
	"todo_rest_api/internal/service"

	"gorm.io/gorm"
)

func buildTodoModule(db *gorm.DB) *handlers.TodoHandler {
	repo := repository.NewTodoRepository(db)
	svc := service.NewTodoService(repo)
	return handlers.NewTodoHandler(svc)
}

func buildUsersModule(db *gorm.DB) *handlers.UsersHandler {
	repo := repository.NewUsersRepository(db)
	svc := service.NewUsersService(repo)
	return handlers.NewUsersHandler(svc)
}

func BuildHandler(db *gorm.DB) *router.Handler {
	return &router.Handler{
		Todo:  buildTodoModule(db),
		Users: buildUsersModule(db),
	}
}
