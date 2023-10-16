package delay

import (
	"context"
	"time"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/google/uuid"
)

type CreateDelayDto struct {
	MessageId uuid.UUID
	Datetime  time.Time
}

type CreateDelayCommand struct {
	delayRepo repository.DelayRepositoryInterface
}

func (cdc *CreateDelayCommand) Exec(ctx context.Context, dto *CreateDelayDto) (*model.Delay, error) {

	delay := model.NewDelay(dto.MessageId)
	delay.DateTime = dto.Datetime

	if err := cdc.delayRepo.Create(ctx, delay); err != nil {
		return nil, err
	}

	return delay, nil
}

func NewCreateDelayCommand(delayRepo repository.DelayRepositoryInterface) *CreateDelayCommand {
	return &CreateDelayCommand{delayRepo: delayRepo}
}
