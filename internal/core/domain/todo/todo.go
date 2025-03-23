package todo

import (
	"time"
	"todo-clean/internal/core/domain"
	"todo-clean/internal/core/enum"
)

type Todo struct {
	Id        domain.ID
	BoardId   domain.ID
	Title     string
	Status    enum.TaskStatus
	Priority  enum.PriorityLevel
	CreatedAt time.Time
	UpdatedAt time.Time
}

var DefaultTodo Todo
