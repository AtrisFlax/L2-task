package entities

import "github.com/google/uuid"

type UserEvents struct {
	UserID uuid.UUID `json:"user_id"`
	Events []Event   `json:"events"`
}
