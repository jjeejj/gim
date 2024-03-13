package repo

import (
	"strings"
	"sync"

	"gim/internal/business/domain/user/model"
)

type userRepo struct {
	sync sync.RWMutex
}

var UserRepo = new(userRepo)

// Get 获取单个用户
func (*userRepo) Get(userId string) (*model.User, error) {
	user, err := UserCache.Get(userId)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = UserDao.GetByUserId(userId)
	if err != nil {
		return nil, err
	}

	if user != nil {
		err = UserCache.Set(*user)
		if err != nil {
			return nil, err
		}
	}
	return user, err
}

func (*userRepo) GetByUserId(userId string) (*model.User, error) {
	return UserDao.GetByUserId(userId)
}

// GetByIds 获取多个用户
func (*userRepo) GetByIds(userIds []string) ([]model.User, error) {
	return UserDao.GetByIds(userIds)
}

// Search 搜索用户
func (*userRepo) Search(key string) ([]model.User, error) {
	return UserDao.Search(key)
}

// Save 保存用户
func (*userRepo) Save(user *model.User) error {
	userId := user.UserId
	err := UserDao.Save(user)
	if err != nil && !strings.Contains(err.Error(), "Duplicate entry") { // TODO 这里要更新 grom 库，进行错误判断
		return err
	}

	if userId != "" {
		err = UserCache.Del(user.UserId)
		if err != nil {
			return err
		}
	}
	return nil
}
