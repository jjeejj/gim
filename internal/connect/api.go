package connect

import (
	"context"
	"fmt"

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

	logger.Logger.Debug("conn info", zap.Any("conn", conn))
	conn.Send(pb.PackageType_PT_MESSAGE, grpclib.GetCtxRequestId(ctx), req.Message, nil)
	return resp, nil
}
