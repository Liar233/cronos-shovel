package message

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/google/uuid"
)

type DeleteMessageDto struct {
	ID uuid.UUID
}

type DeleteMessageCommand struct {
	msgRepo repository.MessageRepositoryInterface
}

func (dmc *DeleteMessageCommand) Exec(ctx context.Context, dto *DeleteMessageDto) error {

	return dmc.msgRepo.Delete(ctx, dto.ID)
}

func NewDeleteMessageCommand(msgRepo repository.MessageRepositoryInterface) *DeleteMessageCommand {
	return &DeleteMessageCommand{msgRepo: msgRepo}
}
