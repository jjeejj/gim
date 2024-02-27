package config

import (
	"context"
	"fmt"

	"google.golang.org/grpc/balancer/roundrobin"

	"gim/pkg/grpclib/picker"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"go.uber.org/zap"

	_ "gim/pkg/grpclib/resolver/addrs"

	"google.golang.org/grpc"
)

type defaultBuilder struct{}

func (*defaultBuilder) Build() Configuration {
	logger.Level = zap.DebugLevel
	logger.Target = logger.Console

	return Configuration{
		MySQL:                "root:Ywkj@2023@tcp(192.168.100.101:3306)/gim?charset=utf8mb4&parseTime=true",
		RedisHost:            "192.168.100.101:3001",
		RedisPassword:        "",
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,

		ConnectLocalAddr:     "192.168.100.101:40000",
		ConnectRPCListenAddr: ":40000",
		ConnectTCPListenAddr: ":40001",
		ConnectWSListenAddr:  ":40002",

		LogicRPCListenAddr:    ":40010",
		BusinessRPCListenAddr: ":40020",
		FileHTTPListenAddr:    "40030",

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///192.168.100.101:40000", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, picker.AddrPickerName)))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///192.168.100.101:40010", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		BusinessIntClientBuilder: func() pb.BusinessIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///192.168.100.101:40020", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewBusinessIntClient(conn)
		},
	}
}
