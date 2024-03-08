package nsq

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"

	"gim/config"
	_const "gim/pkg/const"
	"gim/pkg/logger"
)

var (
	Producer                   *nsq.Producer
	SendMessageChannelConsumer *nsq.Consumer
)

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
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	Producer.Stop()
	SendMessageChannelConsumer.Stop()
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
	// consumer.AddHandler(&sendMessageChannel{})
	// err = consumer.ConnectToNSQD("127.0.0.1:4150")
	err = consumer.ConnectToNSQLookupds(config.Config.Nsq.NsqLookUpdsHost)
	if err != nil {
		logger.Logger.Fatal("consumer ConnectToNSQD err ", zap.Error(err))
	}
	logger.Logger.Info("initNsqSendMessageChannelConsumer success")
	return consumer
}
