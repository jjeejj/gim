package service

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"

	"gim/config"
	"gim/internal/logic/proxy"
	_const "gim/pkg/const"
	"gim/pkg/protocol/pb"

	"gim/pkg/logger"
)

var (
	Producer                   *nsq.Producer
	SendMessageChannelConsumer *nsq.Consumer
)

func init() {
	InitNsq()
}

func InitNsq() {
	logger.Logger.Info("start init nsq")
	Producer = initNsqProducer()
	logger.Logger.Info("Producer", zap.Any("Producer addr", Producer))
	SendMessageChannelConsumer = initNsqSendMessageChannelConsumer()
	logger.Logger.Info("success init nsq")

	// err := Producer.Publish(_const.SEND_MESSAGE_TOPIC_NAME, []byte("hello"))
	// if err != nil {
	//     logger.Logger.Fatal("Publish err ", zap.Error(err))
	// }
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		Producer.Stop()
		SendMessageChannelConsumer.Stop()
	}()

}

// initNsqProducer 初始化生成者
func initNsqProducer() *nsq.Producer {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(config.Config.Nsq.NsqdHost, nsqConfig)
	if err != nil {
		logger.Logger.Fatal("initNsqProducer err ", zap.Error(err))
	}
	logger.Logger.Info("initNsqProducer success", zap.Any("nsq Producer", producer))
	return producer
}

// initNsqSendMessageChannelConsumer 初始化发送消息频道的消费者
func initNsqSendMessageChannelConsumer() *nsq.Consumer {
	nsqConfig := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(_const.SEND_MESSAGE_TOPIC_NAME, _const.SEND_MESSAGE_CHANNEL_NAME, nsqConfig)
	if err != nil {
		logger.Logger.Fatal("initNsqConsumer err ", zap.Error(err))
	}
	logger.Logger.Info("NewConsumer success", zap.Any("Consumer", consumer.Stats()))
	consumer.AddConcurrentHandlers(&sendMessageChannelHandler{}, 1)
	err = consumer.ConnectToNSQLookupds(config.Config.Nsq.NsqLookUpdsHost)
	if err != nil {
		logger.Logger.Fatal("consumer ConnectToNSQD err ", zap.Error(err))
	}
	logger.Logger.Info("initNsqSendMessageChannelConsumer success")
	return consumer
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
