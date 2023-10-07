package message

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
)

type GetMessageCommand struct {
	msgRepo repository.MessageRepositoryInterface
}

func (gmc *GetMessageCommand) Exec(ctx context.Context) ([]*model.Message, error) {

	return gmc.msgRepo.GetList(ctx)
}

func NewGetMessageCommand(msgRepo repository.MessageRepositoryInterface) *GetMessageCommand {
	return &GetMessageCommand{msgRepo: msgRepo}
}
