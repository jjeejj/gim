package repo

import (
	"encoding/json"
	"fmt"
	"strconv"

	"gim/internal/business/domain/user/model"
	"gim/pkg/db"
	"gim/pkg/gerrors"
	"gim/pkg/util"

	"github.com/go-redis/redis"
)

const (
	AuthKey = "auth:"
)

type authCache struct{}

var AuthCache = new(authCache)

func (*authCache) Get(userId string, deviceId int64) (*model.Device, error) {
	bytes, err := db.RedisCli.HGet(fmt.Sprintf("%s%s", AuthKey, userId), strconv.FormatInt(deviceId, 10)).Bytes()
	if err != nil && err != redis.Nil {
		return nil, gerrors.WrapError(err)
	}
	if err == redis.Nil {
		return nil, nil
	}

	var device model.Device
	err = json.Unmarshal(bytes, &device)
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return &device, nil
}

func (*authCache) Set(userId string, deviceId int64, device model.Device) error {
	bytes, err := json.Marshal(device)
	if err != nil {
		return gerrors.WrapError(err)
	}

	_, err = db.RedisCli.HSet(fmt.Sprintf("%s%s", AuthKey, userId), strconv.FormatInt(deviceId, 10), bytes).Result()
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

func (*authCache) GetAll(userId string) (map[int64]model.Device, error) {
	result, err := db.RedisCli.HGetAll(AuthKey + userId).Result()
	if err != nil {
		return nil, gerrors.WrapError(err)
	}

	var devices = make(map[int64]model.Device, len(result))

	for k, v := range result {
		deviceId, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return nil, gerrors.WrapError(err)
		}

		var device model.Device
		err = json.Unmarshal(util.Str2bytes(v), &device)
		if err != nil {
			return nil, gerrors.WrapError(err)
		}
		devices[deviceId] = device
	}
	return devices, nil
}
