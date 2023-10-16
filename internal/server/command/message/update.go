package message

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/google/uuid"
)

type UpdateMessageDto struct {
	ID       uuid.UUID
	Title    string
	Mask     string
	Channels []string
	Payload  []byte
}

type UpdateMessageCommand struct {
	msgRepo repository.MessageRepositoryInterface
}

func (umc *UpdateMessageCommand) Exec(ctx context.Context, dto *UpdateMessageDto) (*model.Message, error) {

	msg := &model.Message{
		ID:       dto.ID,
		Title:    dto.Title,
		Mask:     dto.Mask,
		Channels: dto.Channels,
		Payload:  dto.Payload,
	}

	if err := umc.msgRepo.Update(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func NewUpdateMessageCommand(msgRepo repository.MessageRepositoryInterface) *UpdateMessageCommand {
	return &UpdateMessageCommand{msgRepo: msgRepo}
}
