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

	// ToDo: catch uuid.MustParse panic
	messageId := uuid.MustParse(req.MessageId)

	delay := model.NewDelay(messageId)
	delay.DateTime = time.Unix(req.DateTime, 0)

	resp := &pb.CreateDelayResponse{}

	if err := dc.delayRepo.Create(ctx, delay); err != nil {
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

	// ToDo: catch uuid.MustParse panic
	id := uuid.MustParse(req.Id)

	resp := &pb.DeleteDelayResponse{}

	if err := dc.delayRepo.Delete(ctx, id); err != nil {

		resp.Error = err.Error()
		dc.logger.Error("Create delay error: %s", err.Error())

		return resp, err
	}

	return resp, nil
}

func (dc *DelayController) mustEmbedUnimplementedDelayControllerServer() {
	//TODO implement me
	panic("implement me")
}
