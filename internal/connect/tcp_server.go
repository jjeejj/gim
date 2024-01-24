package connect

import (
	"fmt"
	"time"

	"github.com/alberliu/gn/codec"

	_const "gim/pkg/const"
	"gim/pkg/logger"

	"go.uber.org/zap"

	"github.com/alberliu/gn"
)

var server *gn.Server

// StartTCPServer 启动TCP服务器
func StartTCPServer(addr string) {
	gn.SetLogger(logger.Sugar)

	var err error
	server, err = gn.NewServer(addr, &handler{},
		gn.WithDecoder(codec.NewUvarintDecoder()),
		gn.WithEncoder(codec.NewUvarintEncoder(4098)),
		gn.WithReadBufferLen(4098),
		gn.WithTimeout(11*time.Minute),
		gn.WithAcceptGNum(50),
		gn.WithIOGNum(100))
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
	}
	logger.Logger.Info("tcp server start")
	server.Run()
}

type handler struct{}

func (*handler) OnConnect(c *gn.Conn) {
	// 初始化连接数据
	conn := &Conn{
		CoonType: CoonTypeTCP,
		TCP:      c,
	}
	c.SetData(conn)
	logger.Logger.Debug("connect:", zap.Int32("fd", c.GetFd()), zap.String("addr", c.GetAddr()))
}

func (*handler) OnMessage(c *gn.Conn, bytes []byte) {
	conn := c.GetData().(*Conn)
	conn.HandleMessage(bytes)
}

func (*handler) OnClose(c *gn.Conn, err error) {
	conn, ok := c.GetData().(*Conn)
	if !ok || conn == nil {
		return
	}
	logger.Logger.Debug("close", zap.String("addr", c.GetAddr()), zap.String("user_id", conn.UserId),
		zap.Int64("device_id", conn.DeviceId), zap.Error(err))

	DeleteConn(fmt.Sprintf(_const.CONN_MAP_KEY_FMT, conn.UserId, conn.DeviceId))

	// if conn.UserId != "" {
	// 	_, _ = rpc.GetLogicIntClient().Offline(context.TODO(), &pb.OfflineReq{
	// 		UserId:     conn.UserId,
	// 		DeviceId:   conn.DeviceId,
	// 		ClientAddr: c.GetAddr(),
	// 	})
	// }
}
