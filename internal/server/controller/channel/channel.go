package channel

import (
	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
)

type ChannelInterface interface {
	Fire([]*model.Message) error
	Close() error
}

type ChannelManager struct {
	channels *map[string]ChannelInterface
	msgRepo  repository.MessageRepositoryInterface
}
