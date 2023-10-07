package message

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
)

type CreateMessageDto struct {
	Mask     string
	Title    string
	Channels []string
	Payload  []byte
}

type CreateMessageCommand struct {
	msgRepo repository.MessageRepositoryInterface
}

func (cmc *CreateMessageCommand) Exec(ctx context.Context, dto *CreateMessageDto) error {

	msg := &model.Message{
		Title:    dto.Title,
		Mask:     dto.Mask,
		Channels: dto.Channels,
		Payload:  dto.Payload,
	}

	return cmc.msgRepo.Create(ctx, msg)
}

func NewCreateMessageCommand(msgRepo repository.MessageRepositoryInterface) *CreateMessageCommand {
	return &CreateMessageCommand{msgRepo: msgRepo}
}
