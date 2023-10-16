package message

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
)

type GetMessagesListCommand struct {
	msgRepo repository.MessageRepositoryInterface
}

func (gmc *GetMessagesListCommand) Exec(ctx context.Context) ([]*model.Message, error) {

	return gmc.msgRepo.GetList(ctx)
}

func NewGetMessageCommand(msgRepo repository.MessageRepositoryInterface) *GetMessagesListCommand {
	return &GetMessagesListCommand{msgRepo: msgRepo}
}
