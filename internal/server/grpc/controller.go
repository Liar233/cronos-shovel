package grpc

import (
	"context"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"github.com/sirupsen/logrus"
)

type TikTackController struct {
	logger logrus.FieldLogger
}

func NewTikTackController(logger logrus.FieldLogger) TikTackController {
	return TikTackController{
		logger: logger,
	}
}

func (ttc TikTackController) TikTack(ctx context.Context, req *pb.TikRequest) (*pb.TackResponse, error) {

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		ttc.logger.Info("Request input: %s\n", req.GetInput())
		return &pb.TackResponse{Output: "Tok"}, nil
	}
}
