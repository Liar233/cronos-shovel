package grpc

import (
	"context"
	"time"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type DelayController struct {
	pb.UnimplementedDelayControllerServer
	logger    logrus.FieldLogger
	delayRepo repository.DelayRepositoryInterface
}

func NewDelayController(
	logger logrus.FieldLogger,
	delayRepo repository.DelayRepositoryInterface,
) *DelayController {

	return &DelayController{
		logger:    logger,
		delayRepo: delayRepo,
	}
}

func (dc *DelayController) CreateDelay(ctx context.Context, req *pb.CreateDelayRequest) (*pb.CreateDelayResponse, error) {

	resp := &pb.CreateDelayResponse{}

	messageId, err := uuid.Parse(req.MessageId)

	if err != nil {
		resp.Error = err.Error()
		dc.logger.Error("Invalid delay message_id: %s", err.Error())

		return resp, err
	}

	delay := model.NewDelay(messageId)
	delay.DateTime = time.Unix(req.DateTime, 0)

	if err = dc.delayRepo.Create(ctx, delay); err != nil {
		resp.Error = err.Error()
		dc.logger.Error("Create delay error: %s", err.Error())

		return resp, err
	}

	resp.Id = delay.ID.String()
	resp.MessageId = delay.MessageID.String()
	resp.DateTime = delay.DateTime.Unix()

	return resp, nil
}

func (dc *DelayController) DeleteDelay(ctx context.Context, req *pb.DeleteDelayRequest) (*pb.DeleteDelayResponse, error) {

	resp := &pb.DeleteDelayResponse{}

	id, err := uuid.Parse(req.Id)

	if err != nil {
		resp.Error = err.Error()
		dc.logger.Error("Invalid delay id: %s", err.Error())

		return resp, err
	}

	if err = dc.delayRepo.Delete(ctx, id); err != nil {

		resp.Error = err.Error()
		dc.logger.Error("Delete delay error: %s", err.Error())

		return resp, err
	}

	return resp, nil
}
