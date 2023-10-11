package grpc

import (
	"context"

	pb "github.com/Liar233/cronos-shovel/internal/pkg"
	"github.com/Liar233/cronos-shovel/internal/server/command/message"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type MessageController struct {
	pb.UnimplementedMessageServiceServer
	logger        logrus.FieldLogger
	createMsgCmd  *message.CreateMessageCommand
	updateMsgCmd  *message.UpdateMessageCommand
	deleteMsgCmd  *message.DeleteMessageCommand
	getListMsgCmd *message.GetMessagesListCommand
}

func NewMessageController(
	logger logrus.FieldLogger,
	createMsgCmd *message.CreateMessageCommand,
	updateMsgCmd *message.UpdateMessageCommand,
	deleteMsgCmd *message.DeleteMessageCommand,
	getListMsgCmd *message.GetMessagesListCommand,

) *MessageController {
	return &MessageController{
		logger:        logger,
		createMsgCmd:  createMsgCmd,
		updateMsgCmd:  updateMsgCmd,
		deleteMsgCmd:  deleteMsgCmd,
		getListMsgCmd: getListMsgCmd,
	}
}

func (mc *MessageController) GetMessageList(ctx context.Context, _ *pb.GetMessageListRequest) (*pb.GetMessageListResponse, error) {

	msgList, err := mc.getListMsgCmd.Exec(ctx)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.GetMessageListResponse{}

	for _, msg := range msgList {
		msgObj := &pb.MessageObject{
			Id:       msg.ID.String(),
			Mask:     msg.Mask,
			Title:    msg.Title,
			Channels: msg.Channels,
			Payload:  msg.Payload,
			Delays:   nil,
		}

		resp.Messages = append(resp.Messages, msgObj)
	}

	return resp, nil
}

func (mc *MessageController) CreateMessage(ctx context.Context, req *pb.CreateMessageRequest) (*pb.CreateMessageResponse, error) {

	dto := &message.CreateMessageDto{
		Mask:     req.Mask,
		Title:    req.Mask,
		Channels: req.Channels,
		Payload:  req.Payload,
	}

	msg, err := mc.createMsgCmd.Exec(ctx, dto)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.CreateMessageResponse{
		Id:       msg.ID.String(),
		Mask:     msg.Mask,
		Title:    msg.Title,
		Channels: msg.Channels,
		Payload:  msg.Payload,
	}

	return resp, nil
}

func (mc *MessageController) UpdateMessage(ctx context.Context, req *pb.UpdateMessageRequest) (*pb.UpdateMessageResponse, error) {

	id, err := uuid.Parse(req.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id format")
	}

	dto := &message.UpdateMessageDto{
		ID:       id,
		Title:    req.Title,
		Mask:     req.Mask,
		Channels: req.Channels,
		Payload:  req.Payload,
	}

	msg, err := mc.updateMsgCmd.Exec(ctx, dto)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UpdateMessageResponse{
		Id:       msg.ID.String(),
		Mask:     msg.Mask,
		Title:    msg.Title,
		Channels: msg.Channels,
		Payload:  msg.Payload,
	}

	return resp, nil
}

func (mc *MessageController) DeleteMessage(ctx context.Context, req *pb.DeleteMessageRequest) (*emptypb.Empty, error) {

	id, err := uuid.Parse(req.Id)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id format")
	}

	dto := &message.DeleteMessageDto{ID: id}

	if err = mc.deleteMsgCmd.Exec(ctx, dto); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &emptypb.Empty{}

	return resp, nil
}
