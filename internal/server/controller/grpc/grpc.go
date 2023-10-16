package grpc

import (
	"fmt"
	"net"
	"net/http"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"google.golang.org/grpc"
)

type GRPCConfig struct {
	Port uint32 `yaml:"port"`
	Host string `yaml:"host"`
}

// ToDo: задекорировать как в manners с блокирующей остановкой
type GracefulServer interface {
	http.Handler
	Stop()
	GracefulStop()
}

// NewGRPCServer
// Создание gRPC сервера
func NewGRPCServer(
	config *GRPCConfig,
	msgController pb.MessageServiceServer,
	delayController pb.DelayServiceServer,
) (GracefulServer, error) {

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		return nil, err
	}

	server := grpc.NewServer()
	pb.RegisterMessageServiceServer(server, msgController)
	pb.RegisterDelayServiceServer(server, delayController)

	if err = server.Serve(listener); err != nil {
		return nil, err
	}

	return server, nil
}
