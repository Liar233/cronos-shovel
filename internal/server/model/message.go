package model

import "github.com/google/uuid"

type Message struct {
	ID       uuid.UUID `db:"id"`
	Title    string    `db:"title"`
	Mask     string    `db:"mask"`
	Channels []string  `db:"channels"`
	Payload  []byte    `db:"payload"`
}

func NewMessage() *Message {
	return &Message{
		ID: uuid.New(),
	}
}
