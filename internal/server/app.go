package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Liar233/cronos-shovel/internal/server/command/delay"
	"github.com/Liar233/cronos-shovel/internal/server/command/message"
	"github.com/Liar233/cronos-shovel/internal/server/controller/grpc"
	"github.com/Liar233/cronos-shovel/internal/server/storage"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/sirupsen/logrus"
)

type CronosServerConfig struct {
	GRPC    grpc.GRPCConfig               `yaml:"grpc"`
	Storage storage.PostgresStorageConfig `yaml:"storage"`
}

type CronoServer struct {
	config     *CronosServerConfig
	grpcServer grpc.GracefulServer
	logger     logrus.FieldLogger
	msgRepo    repository.MessageRepositoryInterface
	delayRepo  repository.DelayRepositoryInterface
	storage    storage.ConnectorInterface
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

	cs.storage = storage.NewPostgresStorage(&cs.config.Storage)

	if err := cs.storage.Connect(); err != nil {
		return err
	}

	cs.msgRepo = repository.NewMessagePostgresqlRepository(cs.storage)
	cs.delayRepo = repository.NewDelayPostgresqlRepository(cs.storage)

	return nil
}

func (cs *CronoServer) BootstrapChannels() error {

	return nil
}

func (cs *CronoServer) BootstrapGRPCServer() error {

	createMsgCmd := message.NewCreateMessageCommand(cs.msgRepo)
	updateMsgCmd := message.NewUpdateMessageCommand(cs.msgRepo)
	deleteMsgCmd := message.NewDeleteMessageCommand(cs.msgRepo)
	getMsgListCmd := message.NewGetMessageCommand(cs.msgRepo)

	msgController := grpc.NewMessageController(
		cs.logger,
		createMsgCmd,
		updateMsgCmd,
		deleteMsgCmd,
		getMsgListCmd,
	)

	createDelayCmd := delay.NewCreateDelayCommand(cs.delayRepo)
	deleteDelayCmd := delay.NewDeleteDelayCommand(cs.delayRepo)

	delayController := grpc.NewDelayController(
		cs.logger,
		createDelayCmd,
		deleteDelayCmd,
	)

	var err error

	cs.grpcServer, err = grpc.NewGRPCServer(
		&cs.config.GRPC,
		msgController,
		delayController,
	)

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
