package models

import (
	"time"

	"github.com/google/uuid"
)

type TodoModel struct {
	Id          uuid.UUID `json:"id" gorm:"primaryKey"`
	Task        string    `json:"task"`
	IsCompleted bool      `json:"is_completed"`
	CompletedAt time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
