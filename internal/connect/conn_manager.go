package connect

import (
	"sync"

	"gim/pkg/protocol/pb"
)

// ConnsManager 全局链接 map 维护
// key userid_deviceid
var ConnsManager = sync.Map{}

// SetConn 存储
func SetConn(key string, conn *Conn) {
	ConnsManager.Store(key, conn)
}

// GetConn 根据设备id获取 对应的连接
func GetConn(key string) *Conn {
	value, ok := ConnsManager.Load(key)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(key string) {
	ConnsManager.Delete(key)
}

// PushAll 全服推送
func PushAll(message *pb.Message) {
	ConnsManager.Range(func(key, value interface{}) bool {
		conn := value.(*Conn)
		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)
		return true
	})
}
