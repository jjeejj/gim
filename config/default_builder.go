package config

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"google.golang.org/grpc/balancer/roundrobin"

	"gim/pkg/grpclib/picker"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"go.uber.org/zap"

	_ "gim/pkg/grpclib/resolver/addrs"

	"google.golang.org/grpc"
)

type defaultBuilder struct{}

// loadFileConfig 加载本地配置文件的配置
func loadFileConfig() (*FileConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")  // 编译之后的搜索路径
	viper.AddConfigPath("..") // 开发时候用的搜索路径
	// 加载配置文件
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config file error: %s", err)
	}
	// 解析配置
	var cfg FileConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}
	logger.Logger.Info("load config file success", zap.Any("config", cfg))
	return &cfg, nil
}

func (*defaultBuilder) Build() Configuration {
	logger.Level = zap.InfoLevel
	logger.Target = logger.File

	// 加载配置文件
	cfg, err := loadFileConfig()
	if err != nil {
		logger.Logger.Fatal("load config file error", zap.Error(err))
	}
	return Configuration{
		MySQL:                cfg.MySQL,
		RedisHost:            cfg.RedisHost,
		RedisPassword:        cfg.RedisPassword,
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,

		ConnectLocalAddr:     fmt.Sprintf("%s%s", cfg.ConnectRPCListenHost, cfg.ConnectRPCListenPort),
		ConnectRPCListenAddr: cfg.ConnectRPCListenPort,
		ConnectTCPListenAddr: cfg.ConnectTCPListenAddr,
		// ConnectWSListenAddr:  "172.31.6.248:40002",

		LogicRPCListenAddr:    cfg.LogicRPCListenPort,
		BusinessRPCListenAddr: cfg.BusinessRPCListenPort,
		// FileHTTPListenAddr:    "40030",
		Nsq: cfg.Nsq,

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.DialContext(context.TODO(), fmt.Sprintf("addrs:///%s%s", cfg.ConnectRPCListenHost, cfg.ConnectRPCListenPort), grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, picker.AddrPickerName)))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {

			conn, err := grpc.DialContext(context.TODO(), fmt.Sprintf("addrs:///%s%s", cfg.LogicRPCListenHost, cfg.LogicRPCListenPort), grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		BusinessIntClientBuilder: func() pb.BusinessIntClient {
			conn, err := grpc.DialContext(context.TODO(), fmt.Sprintf("addrs:///%s%s", cfg.BusinessRPCListenHost, cfg.BusinessRPCListenPort), grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewBusinessIntClient(conn)
		},
	}
}
