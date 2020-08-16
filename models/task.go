package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Task struct {
	ID uuid.UUID `json:"id" db:"id"`

	Description    string    `json:"description"`
	Status         string    `json:"status"`
	CompletionDate time.Time `json:"completion_date"`
	Requester      string    `json:"requester"`
	Executor       string    `json:"executor"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
