package router

import (
	"todo_rest_api/internal/handlers"
	"todo_rest_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Todo  *handlers.TodoHandler
	Users *handlers.UsersHandler
}

func Setup(h *Handler) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message":  "TODOS RestAPI",
			"status":   "200",
			"database": "connected",
		})
	})

	api := r.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("", h.Users.Create)
			users.POST("/login", h.Users.Login)
		}
		todos := api.Group("/todos")
		todos.Use(middleware.AuthRequired())
		{
			todos.POST("", h.Todo.Create)
			todos.GET("", h.Todo.GetAll)
			todos.GET("/:id", h.Todo.GetByID)
			todos.PUT("/:id", h.Todo.Update)
			todos.DELETE("/:id", h.Todo.Delete)
		}
	}
	return r
}
