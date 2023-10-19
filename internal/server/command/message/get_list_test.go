package message

import (
	"context"
	"reflect"
	"testing"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	mocks "github.com/Liar233/cronos-shovel/internal/server/storage/repository/mocks"
)

func TestGetMessagesListCommand_Exec(t *testing.T) {
	ctx := context.Background()

	mockMessageRepository := mocks.NewMessageMockRepository(t)
	mockMessageRepository.
		On("GetList", ctx).
		Once().
		Return(make([]*model.Message, 0), nil)

	getListCmd := NewGetMessageCommand(mockMessageRepository)
	messages, err := getListCmd.Exec(ctx)

	if err != nil {
		t.Fatalf("fail creget message list with: %s", err.Error())
	}

	println(reflect.TypeOf(messages))
}
