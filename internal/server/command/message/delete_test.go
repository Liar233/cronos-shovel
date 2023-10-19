package message

import (
	"context"
	"testing"

	mocks "github.com/Liar233/cronos-shovel/internal/server/storage/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var messageUUID uuid.UUID = uuid.MustParse("40e9d86f-b285-4d57-bcd0-5731e0cf7595")

func TestDeleteMessageCommand_Exec(t *testing.T) {

	dto := buildDeleteDto()
	ctx := context.Background()

	mockMessageRepository := mocks.NewMessageMockRepository(t)
	mockMessageRepository.
		On(
			"Delete",
			ctx,
			mock.MatchedBy(func(val interface{}) bool {
				_, ok := val.(uuid.UUID)
				return ok
			}),
		).
		Once().
		Return(nil)

	deleteCmd := NewDeleteMessageCommand(mockMessageRepository)
	err := deleteCmd.Exec(ctx, dto)

	if err != nil {
		t.Fatalf("fail delete message with: %s", err.Error())
	}
}

func buildDeleteDto() *DeleteMessageDto {

	return &DeleteMessageDto{
		ID: messageUUID,
	}
}
