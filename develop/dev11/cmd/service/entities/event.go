package entities

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Date   time.Time `json:"date"`
	UserID uuid.UUID `json:"-"`
	ID     uuid.UUID `json:"event_id"`
	Info   string    `json:"info"`
}
