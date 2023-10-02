package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Liar233/cronos-shovel/internal/server/grpc"
	"github.com/sirupsen/logrus"
)

type CronoServer struct {
	config     *CronosServerConfig
	grpcServer grpc.GracefulServer
	logger     logrus.FieldLogger
}

func NewCronoServer(config *CronosServerConfig) *CronoServer {
	return &CronoServer{
		config: config,
	}
}

func (cs *CronoServer) Run() {

	_ = cs.BootstrapLogger()

	cs.logger.Info("Starting Cronos' Shovel grpc")

	cs.logger.Debug("Storage bootstrapping")
	if err := cs.BootstrapStorage(); err != nil {
		cs.logger.Errorf("Storage error: %s", err.Error())
		return
	}
	cs.logger.Debug("Storage bootstrapped")

	cs.logger.Debug("Channels bootstrapping")
	if err := cs.BootstrapChannels(); err != nil {
		cs.logger.Errorf("Channels error: %s", err.Error())
		return
	}
	cs.logger.Debug("Channels bootstrapped")

	cs.logger.Debug("GRPC-grpc bootstrapping")
	if err := cs.BootstrapGRPCServer(); err != nil {
		cs.logger.Errorf("GRPC-grpc error: %s", err.Error())
		return
	}
	cs.logger.Debug("GRPC-grpc bootstrapped")

	cs.logger.Info("Cronos' Shovel grpc started")

	if err := cs.handleClose(); err != nil {
		cs.logger.Errorf("Stopping Cronos' Shovel grpc error: %s", err.Error())
	}
}

func (cs *CronoServer) BootstrapStorage() error {

	return nil
}

func (cs *CronoServer) BootstrapChannels() error {

	return nil
}

func (cs *CronoServer) BootstrapGRPCServer() error {
	var err error

	controller := grpc.NewTikTackController(cs.logger)

	cs.grpcServer, err = grpc.NewGRPCServer(&cs.config.GRPC, controller)

	return err
}

func (cs *CronoServer) handleClose() error {

	sigintChan := make(chan os.Signal, 1)

	defer close(sigintChan)

	signal.Notify(sigintChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigintChan

	cs.logger.Info("Cronos' Shovel grpc got signal %s", sig)

	// ToDo: implement graceful stop for GRPCserver

	// ToDo: implement graceful stop for Channels

	// ToDo: implement graceful stop for Storage

	return nil
}

func (cs *CronoServer) BootstrapLogger() error {
	logger := logrus.New()

	logger.Out = os.Stdout
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.DebugLevel)

	cs.logger = logger

	return nil
}

type CronosServerConfig struct {
	GRPC grpc.GRPCConfig `yaml:"grpc"`
}
