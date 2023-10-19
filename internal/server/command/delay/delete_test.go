package delay

import (
	"context"
	"testing"

	mocks "github.com/Liar233/cronos-shovel/internal/server/storage/repository/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

var delayUUID uuid.UUID = uuid.MustParse("40e9d86f-b285-4d57-bcd0-5731e0cf7595")

func TestDeleteDelayCommand_Exec(t *testing.T) {

	dto := buildDeleteDto()
	ctx := context.Background()

	mockDelayRepository := mocks.NewDelayMockRepository(t)
	mockDelayRepository.
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

	deleteCmd := NewDeleteDelayCommand(mockDelayRepository)
	err := deleteCmd.Exec(ctx, dto)

	if err != nil {
		t.Fatalf("delete delay with: %s", err.Error())
	}
}

func buildDeleteDto() *DeleteDelayDto {

	return &DeleteDelayDto{
		ID: delayUUID,
	}
}
