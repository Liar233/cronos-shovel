package message

import (
	"context"
	"testing"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	mocks "github.com/Liar233/cronos-shovel/internal/server/storage/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func TestCreateMessageCommand_Exec(t *testing.T) {

	dto := buildCreateDto()
	ctx := context.Background()

	mockMessageRepository := mocks.NewMessageMockRepository(t)
	mockMessageRepository.
		On(
			"Create",
			ctx,
			mock.MatchedBy(func(val interface{}) bool {
				_, ok := val.(*model.Message)
				return ok
			}),
		).
		Once().
		Return(nil)

	createCmd := NewCreateMessageCommand(mockMessageRepository)
	msgModel, err := createCmd.Exec(ctx, dto)

	if err != nil {
		t.Fatalf("fail create message with: %s", err.Error())
	}

	if msgModel.ID == uuid.Nil {
		t.Fatalf("empty message id")
	}
}

func buildCreateDto() *CreateMessageDto {

	return &CreateMessageDto{
		Mask:     "* * * * *",
		Title:    "Test",
		Channels: []string{"test_chan"},
		Payload:  []byte("Zii"),
	}
}
