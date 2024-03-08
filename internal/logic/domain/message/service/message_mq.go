package service

import (
	"context"
	"encoding/json"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"

	logicNsq "gim/internal/logic/nsq"
	"gim/internal/logic/proxy"
	"gim/pkg/protocol/pb"

	"gim/pkg/logger"
)

func init() {
	logicNsq.SendMessageChannelConsumer.AddConcurrentHandlers(&sendMessageChannelHandler{}, 1)
}

// NsgMessage 发送到 nsq 队列中的 消息格式
type NsgMessage struct {
	message  *pb.Message
	toUserID string
}

type sendMessageChannelHandler struct{}

// HandleMessage 异步处理发送消息
func (c *sendMessageChannelHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}

	// 打印接受到的消息
	logger.Logger.Info("receive message", zap.String("message", string(m.Body)))

	nsgMessage := &NsgMessage{}
	err := json.Unmarshal(m.Body, nsgMessage)
	if err != nil {
		logger.Logger.Error("Unmarshal Body message", zap.Error(err))
	}
	ctx := context.Background()
	// 查询用户在线设备
	devices, err := proxy.DeviceProxy.ListOnlineByUserId(ctx, nsgMessage.toUserID)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	for i := range devices {
		// 消息不需要投递给发送消息的设备, 自己发送的也要推送
		// if fromDeviceID == devices[i].DeviceId {
		//     continue
		// }
		err = MessageService.SendToDevice(ctx, devices[i], nsgMessage.toUserID, nsgMessage.message)
		if err != nil {
			logger.Sugar.Error(err, zap.Any("SendToUser error", devices[i]), zap.Error(err))
		}
	}
	return nil
}
