package repository

import (
	"context"

	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/google/uuid"
)

type MessagePostgresqlRepository struct {
	dialect  goqu.DialectWrapper
	provider storage.ConnectorInterface
}

func NewMessagePostgresqlRepository(provider storage.ConnectorInterface) *MessagePostgresqlRepository {

	return &MessagePostgresqlRepository{
		dialect:  goqu.Dialect("postgres"),
		provider: provider,
	}
}

func (mpr *MessagePostgresqlRepository) GetList(ctx context.Context) ([]*model.Message, error) {
	var messages []*model.Message

	err := mpr.dialect.
		From("messages").
		LeftJoin(
			goqu.T("dalays"),
			goqu.On(
				goqu.Ex{"messages.id": goqu.I("delays.message_id")},
			),
		).
		ScanStructsContext(ctx, messages)

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (mpr *MessagePostgresqlRepository) Create(ctx context.Context, msg *model.Message) error {
	_, err := mpr.dialect.
		Insert("messages").
		Rows(msg).
		Executor().
		ExecContext(ctx)

	return err
}

func (mpr *MessagePostgresqlRepository) Update(ctx context.Context, msg *model.Message) error {
	_, err := mpr.dialect.
		Update("messages").
		Set(msg).
		Executor().
		ExecContext(ctx)

	return err
}

func (mpr *MessagePostgresqlRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := mpr.dialect.
		Delete("messages").
		Where(
			goqu.C("id").Eq(id),
		).
		Executor().
		ExecContext(ctx)

	return err
}

func (mpr *MessagePostgresqlRepository) FindOne(ctx context.Context, id uuid.UUID) (*model.Message, error) {
	var msg *model.Message

	_, err := mpr.dialect.
		From("messages").
		LeftJoin(
			goqu.T("delays"),
			goqu.On(
				goqu.Ex{"messages.id": goqu.I("delays.message_id")},
			),
		).
		Where(goqu.C("messages.id").Eq(id)).
		ScanStructContext(ctx, msg)

	if err != nil {
		return nil, err
	}

	return msg, nil
}
