package delay

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/google/uuid"
)

type DeleteDelayDto struct {
	ID uuid.UUID
}

type DeleteDelayCommand struct {
	delayRepo repository.DelayRepositoryInterface
}

func (ddc *DeleteDelayCommand) Exec(ctx context.Context, dto *DeleteDelayDto) error {

	return ddc.delayRepo.Delete(ctx, dto.ID)
}

func NewDeleteDelayCommand(delayRepo repository.DelayRepositoryInterface) *DeleteDelayCommand {
	return &DeleteDelayCommand{delayRepo: delayRepo}
}
