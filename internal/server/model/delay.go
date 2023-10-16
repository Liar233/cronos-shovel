package model

import (
	"time"

	"github.com/google/uuid"
)

type Delay struct {
	ID        uuid.UUID `db:"id"`
	MessageID uuid.UUID `db:"message_id"`
	DateTime  time.Time `db:"datetime"`
}

func NewDelay(messageID uuid.UUID) *Delay {
	return &Delay{
		ID:        uuid.New(),
		MessageID: messageID,
	}
}
