package repo

import (
	"time"

	"gim/internal/logic/domain/group/entity"
	"gim/pkg/db"
	"gim/pkg/gerrors"

	"github.com/go-redis/redis"
)

const GroupKey = "group:"

type groupCache struct{}

var GroupCache = new(groupCache)

// Get 获取群组缓存
func (c *groupCache) Get(groupId string) (*entity.Group, error) {
	var group entity.Group
	err := db.RedisUtil.Get(GroupKey+groupId, &group)
	if err != nil && err != redis.Nil {
		return nil, gerrors.WrapError(err)
	}
	if err == redis.Nil {
		return nil, nil
	}
	return &group, nil
}

// Set 设置群组缓存
// 有效期为 24h
func (c *groupCache) Set(group *entity.Group) error {
	err := db.RedisUtil.Set(GroupKey+group.GroupId, group, 24*time.Hour)
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// Del 删除群组缓存
func (c *groupCache) Del(groupId string) error {
	_, err := db.RedisCli.Del(GroupKey + groupId).Result()
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}
