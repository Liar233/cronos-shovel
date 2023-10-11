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

func (cmc *CreateMessageCommand) Exec(ctx context.Context, dto *CreateMessageDto) (*model.Message, error) {

	msg := model.NewMessage()
	msg.Title = dto.Title
	msg.Mask = dto.Mask
	msg.Channels = dto.Channels
	msg.Payload = dto.Payload

	if err := cmc.msgRepo.Create(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func NewCreateMessageCommand(msgRepo repository.MessageRepositoryInterface) *CreateMessageCommand {
	return &CreateMessageCommand{msgRepo: msgRepo}
}
