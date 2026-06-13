package dto

import "time"

type CreateTodoReq struct {
	Title     string `json:"title" binding:"required"`
	Completed bool   `json:"completed"`
}

type UpdateTodoReq struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

type TodoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
