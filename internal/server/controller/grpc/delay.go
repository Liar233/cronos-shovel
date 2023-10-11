package grpc

import (
	"context"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"github.com/Liar233/cronos-shovel/internal/server/command/delay"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DelayController struct {
	pb.UnimplementedDelayServiceServer
	logger         logrus.FieldLogger
	createDelayCmd *delay.CreateDelayCommand
	deleteDelayCmd *delay.DeleteDelayCommand
}

func NewDelayController(
	logger logrus.FieldLogger,
	createDelayCmd *delay.CreateDelayCommand,
	deleteDelayCmd *delay.DeleteDelayCommand,
) *DelayController {

	return &DelayController{
		logger:         logger,
		createDelayCmd: createDelayCmd,
		deleteDelayCmd: deleteDelayCmd,
	}
}

func (dc *DelayController) CreateDelay(ctx context.Context, req *pb.CreateDelayRequest) (*pb.CreateDelayResponse, error) {

	msgId, err := uuid.Parse(req.MessageId)

	if err != nil {
		return nil, err
	}

	dto := &delay.CreateDelayDto{
		MessageId: msgId,
		Datetime:  req.DateTime.AsTime(),
	}

	delay, err := dc.createDelayCmd.Exec(ctx, dto)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	ts := timestamppb.New(delay.DateTime)

	resp := &pb.CreateDelayResponse{
		Id:        delay.ID.String(),
		MessageId: delay.MessageID.String(),
		DateTime:  ts,
	}

	return resp, nil
}

func (dc *DelayController) DeleteDelay(ctx context.Context, req *pb.DeleteDelayRequest) (*emptypb.Empty, error) {

	id, err := uuid.Parse(req.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid delay id")
	}

	dto := &delay.DeleteDelayDto{ID: id}

	if err = dc.deleteDelayCmd.Exec(ctx, dto); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
