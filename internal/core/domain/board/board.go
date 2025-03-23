package board

import (
	"time"
	"todo-clean/internal/core/domain"
	"todo-clean/internal/core/enum"
)

type Board struct {
	Id        domain.ID
	Name      string
	status    enum.BoardStatus
	TaskIds   []domain.ID
	CreatedAt time.Time
}

var DefaultBoard Board
