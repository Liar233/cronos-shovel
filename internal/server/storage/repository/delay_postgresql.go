package repository

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type DelayPostgresqlRepository struct {
	dialect  goqu.DialectWrapper
	provider storage.ConnectorInterface
}

func NewDelayPostgresqlRepository(provider storage.ConnectorInterface) *DelayPostgresqlRepository {
	return &DelayPostgresqlRepository{
		dialect:  goqu.Dialect("postgres"),
		provider: provider,
	}
}

func (dpr *DelayPostgresqlRepository) Create(ctx context.Context, delay *model.Delay) error {
	//TODO implement me
	panic("implement me")
}

func (dpr *DelayPostgresqlRepository) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
