package models

import "time"

type Todos struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Title     string    `json:"title" gorm:"column:title"`
	Completed bool      `json:"completed" gorm:"column:completed"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (Todos) TableName() string {
	return "golang.todos_user"
}
