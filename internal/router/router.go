package router

import (
	"todo_rest_api/internal/handlers"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Todo *handlers.TodoHandler
}

func Setup(h *TodoHandler) *gin.Engine {
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
		todos := api.Group("/todos")
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
