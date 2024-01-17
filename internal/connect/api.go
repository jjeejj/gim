package connect

import (
	"context"

	"github.com/bytedance/sonic"

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
	conn := GetConn(req.DeviceId)
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
	// 反序列化，填充 社交id
	userMessagePus := &pb.UserMessagePush{}
	err := sonic.Unmarshal(req.Message.Content, userMessagePus)
	if err == nil {
		msgContent := &pb.GimMessage{}
		err = sonic.Unmarshal(userMessagePus.Content, msgContent)
		if err == nil {
			msgContent.SocialMsgId = req.Message.UserSeq
			req.Message.Content, _ = sonic.Marshal(userMessagePus)
		}
	}
	conn.Send(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequestId(ctx), req.Message, nil)
	return resp, nil
}
