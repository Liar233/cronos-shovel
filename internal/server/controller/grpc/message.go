package grpc

import (
	"context"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"github.com/Liar233/cronos-shovel/internal/server/model"
	"github.com/Liar233/cronos-shovel/internal/server/storage/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type MessageController struct {
	pb.UnimplementedMessageControllerServer
	logger  logrus.FieldLogger
	msgRepo repository.MessageRepositoryInterface
}

func NewMessageController(
	logger logrus.FieldLogger,
	msgRepo repository.MessageRepositoryInterface,
) *MessageController {

	return &MessageController{
		logger:  logger,
		msgRepo: msgRepo,
	}
}

func (mc *MessageController) GetMessageList(ctx context.Context, _ *pb.GetMessageListRequest) (*pb.GetMessageListResponse, error) {

	resp := &pb.GetMessageListResponse{}

	msgList, err := mc.msgRepo.GetList(ctx)

	if err != nil {
		resp.Error = err.Error()
		mc.logger.Error("Get messages list error: %s", err.Error())

		return resp, err
	}

	for _, msg := range msgList {

		obj := &pb.MessageObject{
			Id:       msg.ID.String(),
			Mask:     msg.Mask,
			Title:    msg.Title,
			Channels: msg.Channels,
			Payload:  msg.Payload,
			Delays:   nil,
		}

		resp.Messages = append(resp.Messages, obj)
	}

	return resp, err
}

func (mc *MessageController) CreateMessage(ctx context.Context, req *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {

	msg := model.NewMessage()
	msg.Mask = req.Mask
	msg.Title = req.Title
	msg.Channels = req.Channels
	msg.Payload = req.Payload

	err := mc.msgRepo.Create(ctx, msg)

	resp := &pb.CreateMessageResponse{}

	if err != nil {
		resp.Error = err.Error()
		mc.logger.Error("Create message error: %s", err.Error())

		return resp, err
	}

	resp.Id = msg.ID.String()
	resp.Title = msg.Title
	resp.Mask = msg.Mask
	resp.Channels = msg.Channels
	resp.Payload = msg.Payload

	return resp, nil
}

func (mc *MessageController) UpdateMessage(ctx context.Context, req *pb.UpdateMessageRequest) (*pb.UpdateMessageResponse, error) {

	resp := &pb.UpdateMessageResponse{}

	id, err := uuid.Parse(req.Id)

	if err != nil {
		resp.Error = err.Error()
		mc.logger.Error("Invalid message id: %s", err.Error())

		return resp, err
	}

	msg := &model.Message{
		ID:       id,
		Title:    req.Title,
		Mask:     req.Mask,
		Channels: req.Channels,
		Payload:  req.Payload,
	}

	err = mc.msgRepo.Update(ctx, msg)

	if err != nil {
		resp.Error = err.Error()
		mc.logger.Error("Update message error: %s", err.Error())

		return resp, err
	}

	resp.Id = msg.ID.String()
	resp.Title = msg.Title
	resp.Mask = msg.Mask
	resp.Channels = msg.Channels
	resp.Payload = msg.Payload

	return resp, nil
}

func (mc *MessageController) DeleteMessage(ctx context.Context, req *pb.DeleteMessageRequest) (*pb.DeleteMessageResponse, error) {

	resp := &pb.DeleteMessageResponse{}

	id, err := uuid.Parse(req.Id)

	if err != nil {
		resp.Error = err.Error()
		mc.logger.Error("Invalid message id: %s", err.Error())

		return resp, err
	}

	if err = mc.msgRepo.Delete(ctx, id); err != nil {
		resp.Error = err.Error()
		mc.logger.Error("Delete message error: %s", err.Error())

		return resp, err
	}

	return resp, nil
}
