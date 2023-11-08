package repo

import (
	"time"

	"gim/internal/business/domain/user/model"
	"gim/pkg/db"
	"gim/pkg/gerrors"

	"github.com/jinzhu/gorm"
)

type userDao struct{}

var UserDao = new(userDao)

// Add 插入一条用户信息
func (*userDao) Add(user model.User) (int64, error) {
	user.CreateTime = time.Now()
	user.UpdateTime = time.Now()
	err := db.DB.Create(&user).Error
	if err != nil {
		return 0, gerrors.WrapError(err)
	}
	return user.Id, nil
}

// Get 获取用户信息
func (*userDao) Get(userId int64) (*model.User, error) {
	var user = model.User{Id: userId}
	err := db.DB.First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, gerrors.WrapError(err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

// Save 保存
func (*userDao) Save(user *model.User) error {
	err := db.DB.Save(user).Error
	if err != nil {
		return gerrors.WrapError(err)
	}
	return nil
}

// GetByUserId 根据用户业务唯一id获取用户信息
func (*userDao) GetByUserId(userId string) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, "user_id = ?", userId).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, gerrors.WrapError(err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
}

// GetByIds 获取用户信息
func (*userDao) GetByIds(userIds []string) ([]model.User, error) {
	var users []model.User
	err := db.DB.Find(&users, "user_id in (?)", userIds).Error
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return users, err
}

// Search 查询用户,这里简单实现，生产环境建议使用ES
func (*userDao) Search(key string) ([]model.User, error) {
	var users []model.User
	key = "%" + key + "%"
	err := db.DB.Where("user_id like ? or nickname like ?", key, key).Find(&users).Error
	if err != nil {
		return nil, gerrors.WrapError(err)
	}
	return users, nil
}
