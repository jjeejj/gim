package device

import (
	"time"

	"gim/pkg/db"
	"gim/pkg/gerrors"

	"github.com/go-redis/redis"
)

const (
	UserDeviceKey    = "user_device:"
	UserDeviceExpire = 2 * time.Hour
)

type userDeviceCache struct{}

var UserDeviceCache = new(userDeviceCache)

// Get 获取指定用户的所有在线设备
func (c *userDeviceCache) Get(userId string) ([]Device, error) {
	var devices []Device
	err := db.RedisUtil.Get(UserDeviceKey+userId, &devices)
	if err != nil && err != redis.Nil {
		return nil, gerrors.WrapError(err)
	}

	if err == redis.Nil {
		return nil, nil
	}
	return devices, nil
}

// Set 将指定用户的所有在线设备存入缓存
func (c *userDeviceCache) Set(userId string, devices []Device) error {
	err := db.RedisUtil.Set(UserDeviceKey+userId, devices, UserDeviceExpire)
	return gerrors.WrapError(err)
}

// Del 删除用户的在线设备列表
func (c *userDeviceCache) Del(userId string) error {
	key := UserDeviceKey + userId
	_, err := db.RedisCli.Del(key).Result()
	return gerrors.WrapError(err)
}
