package delay

import (
	"context"
	"testing"
	"time"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	mocks "github.com/Liar233/cronos-shovel/internal/server/storage/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var messageUUID uuid.UUID = uuid.MustParse("40e9d86f-b285-4d57-bcd0-5731e0cf7595")

func TestCreateDelayCommand_Exec(t *testing.T) {

	dto := buildCreateDto()
	ctx := context.Background()

	mockDelayRepository := mocks.NewDelayMockRepository(t)
	mockDelayRepository.
		On(
			"Create",
			ctx,
			mock.MatchedBy(func(val interface{}) bool {
				_, ok := val.(*model.Delay)
				return ok
			}),
		).
		Once().
		Return(nil)

	createCmd := NewCreateDelayCommand(mockDelayRepository)
	delayModel, err := createCmd.Exec(ctx, dto)

	if err != nil {
		t.Fatalf("fail create delay with: %s", err.Error())
	}

	if delayModel.MessageID != messageUUID {
		t.Fatalf("invalid message_id while creating delay")
	}

	if delayModel.ID == uuid.Nil {
		t.Fatalf("empty delay id")
	}
}

func buildCreateDto() *CreateDelayDto {

	return &CreateDelayDto{
		MessageId: messageUUID,
		Datetime:  time.Time{},
	}
}
