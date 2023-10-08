package repository

import (
	"context"
	"errors"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/google/uuid"
)

var (
	InvalidMessageModelError  = errors.New("invalid message data")
	MessageAlreadyExistsError = errors.New("message already exists")
	MessageDoesNotExistError  = errors.New("message does not exists")
)

type MessageRepositoryInterface interface {
	GetList(ctx context.Context) ([]*model.Message, error)
	Create(ctx context.Context, msg *model.Message) error
	Update(ctx context.Context, msg *model.Message) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindOne(ctx context.Context, id uuid.UUID) (*model.Message, error)
}

type DelayRepositoryInterface interface {
	Create(ctx context.Context, msg *model.Delay) error
	Delete(ctx context.Context, id uuid.UUID) error
}
