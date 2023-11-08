package repo

import (
	"fmt"
	"strconv"
	"time"

	"gim/internal/business/domain/user/model"
	"gim/pkg/db"
	"gim/pkg/gerrors"

	"github.com/go-redis/redis"
)

const (
	UserKey    = "user:"
	UserExpire = 2 * time.Hour
)

type userCache struct{}

var UserCache = new(userCache)

// Get 获取用户缓存
func (c *userCache) Get(userId string) (*model.User, error) {
	var user model.User
	err := db.RedisUtil.Get(fmt.Sprintf("%s%s", UserKey, userId), &user)
	if err != nil && err != redis.Nil {
		return nil, gerrors.WrapError(err)
	}
	if err == redis.Nil {
		return nil, nil
	}
	return &user, nil
}

// Set 设置用户缓存
func (c *userCache) Set(user model.User) error {
	err := db.RedisUtil.Set(UserKey+strconv.FormatInt(user.Id, 10), user, UserExpire)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// Del 删除用户缓存
func (c *userCache) Del(userId string) error {
	_, err := db.RedisCli.Del(fmt.Sprintf("%s%s", UserKey, userId)).Result()
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
