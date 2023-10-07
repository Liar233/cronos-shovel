package repository

import (
	"sync"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/google/uuid"
)

type DelayMemoryRepository struct {
	storage sync.Map
}

func (dmr *DelayMemoryRepository) Save(delay *model.Delay) error {

	dmr.storage.Store(delay.ID.String(), delay)

	return nil
}

func (dmr *DelayMemoryRepository) Delete(ID uuid.UUID) error {

	dmr.storage.Delete(ID.String())

	return nil
}

func NewDelayMemoryRepository() *DelayMemoryRepository {
	return &DelayMemoryRepository{}
}
