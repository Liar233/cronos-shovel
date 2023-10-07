package repository

import (
	"sync"

	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/google/uuid"
)

type MessageMemoryRepository struct {
	storage sync.Map
	count   uint16
}

func NewMessageMemoryRepository() *MessageMemoryRepository {
	return &MessageMemoryRepository{}
}

func (mmr *MessageMemoryRepository) GetList(ctx context.Context) ([]*model.Message, error) {

	result := make([]*model.Message, mmr.count)

	i := 0
	mmr.storage.Range(func(key, value interface{}) bool {
		result[i] = value.(*model.Message)
		i++

		return true
	})

	return result, nil
}

func (mmr *MessageMemoryRepository) Create(ctx context.Context, message *model.Message) error {

	if message.ID == uuid.Nil {
		return InvalidMessageModelError
	}

	if _, ok := mmr.storage.Load(message.ID.String()); ok {
		return MessageAlreadyExistsError
	}

	mmr.storage.Store(message.ID.String(), message)

	return nil
}

func (mmr *MessageMemoryRepository) Update(ctx context.Context, message *model.Message) error {

	if message.ID == uuid.Nil {
		return InvalidMessageModelError
	}

	if _, ok := mmr.storage.Load(message.ID.String()); !ok {
		return MessageDoesNotExistError
	}

	mmr.storage.Store(message.ID.String(), message)

	return nil
}

func (mmr *MessageMemoryRepository) Delete(ctx context.Context, ID uuid.UUID) error {

	if _, ok := mmr.storage.Load(ID.String()); !ok {
		return MessageDoesNotExistError
	}

	mmr.storage.Delete(ID.String())

	return nil
}

func (mmr *MessageMemoryRepository) FindOne(ctx context.Context, ID uuid.UUID) (*model.Message, error) {

	if item, ok := mmr.storage.Load(ID.String()); ok {

		return item.(*model.Message), nil
	}

	return nil, MessageDoesNotExistError
}
