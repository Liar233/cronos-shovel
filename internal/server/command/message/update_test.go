package message

import (
	"context"
	"testing"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	mocks "github.com/Liar233/cronos-shovel/internal/server/storage/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestUpdateMessageCommand_Exec(t *testing.T) {

	dto := buildUpdateDto()
	ctx := context.Background()

	mockMessageRepository := mocks.NewMessageMockRepository(t)
	mockMessageRepository.
		On(
			"Update",
			ctx,
			mock.MatchedBy(func(val interface{}) bool {
				_, ok := val.(*model.Message)
				return ok
			}),
		).
		Once().
		Return(nil)

	updateCmd := NewUpdateMessageCommand(mockMessageRepository)
	msgModel, err := updateCmd.Exec(ctx, dto)

	if err != nil {
		t.Fatalf("fail updaate message with: %s", err.Error())
	}

	if msgModel.ID == uuid.Nil {
		t.Fatalf("empty message id")
	}
}

func buildUpdateDto() *UpdateMessageDto {

	return &UpdateMessageDto{
		ID:       messageUUID,
		Mask:     "* * * * *",
		Title:    "Test",
		Channels: []string{"test_chan"},
		Payload:  []byte("Zii"),
	}
}
