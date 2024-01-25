package connect

import (
	"context"
	"fmt"

	"github.com/bytedance/sonic"
	"google.golang.org/protobuf/proto"

	_const "gim/pkg/const"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/zap"
)

type ConnIntServer struct {
	pb.UnsafeConnectIntServer
}

// DeliverMessage 投递消息
func (s *ConnIntServer) DeliverMessage(ctx context.Context, req *pb.DeliverMessageReq) (*emptypb.Empty, error) {
	resp := &emptypb.Empty{}
	// 获取设备对应的TCP连接
	conn := GetConn(fmt.Sprintf(_const.CONN_MAP_KEY_FMT, req.UserId, req.DeviceId))
	if conn == nil {
		logger.Logger.Warn("conn is nil GetConn warn", zap.Int64("device_id", req.DeviceId))
		return resp, nil
	}
	if conn.DeviceId != req.DeviceId {
		logger.Logger.Warn("conn.DeviceId is not equal GetConn req.DeviceId warn",
			zap.Int64("device_id", req.DeviceId),
			zap.Int64("conn_device_id", conn.DeviceId))
		return resp, nil
	}
	logger.Logger.Warn("req.Message", zap.Any("req.Message", req.Message))
	// 反序列化，填充 社交id
	userMessagePus := &pb.UserMessagePush{}
	err := proto.Unmarshal(req.Message.Content, userMessagePus)
	if err != nil {
		logger.Logger.Warn("sonic.Unmarshal userMessagePus error", zap.Error(err))
		// return resp, nil
	} else {
		// logger.Logger.Info("sonic.Unmarshal userMessagePus success", zap.Any("userMessagePus", userMessagePus))
		msgContent := &pb.GimMessage{}
		err = sonic.Unmarshal(userMessagePus.Content, msgContent)
		if err != nil {
			logger.Logger.Warn("sonic.Unmarshal msgContent error", zap.Error(err))
		} else {
			msgContent.SocialMsgId = req.Message.UserSeq
			userMessagePus.Content, _ = sonic.Marshal(msgContent)
			req.Message.Content, _ = proto.Marshal(userMessagePus)
			// logger.Logger.Info("sonic.Unmarshal msgContent success", zap.Any("msgContent", msgContent))
		}
	}
	logger.Logger.Debug("conn info", zap.Any("conn", conn))
	conn.Send(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequestId(ctx), req.Message, nil)
	return resp, nil
}
