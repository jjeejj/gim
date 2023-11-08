package service

import (
	"context"
	"time"

	"gim/internal/business/domain/user/model"
	"gim/internal/business/domain/user/repo"
	"gim/pkg/gerrors"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"
	"gim/pkg/util"
)

type authService struct{}

var AuthService = new(authService)

// SignIn 登录
// 判断是否是新用户，记录用户信息
// 判断设备是否存在
// 生成 token
// 记录到 redis 缓存中
func (*authService) SignIn(ctx context.Context, userId string, deviceId int64, sourceCode string) (bool, string, error) {

	user, err := repo.UserRepo.GetByUserId(userId)
	if err != nil {
		return false, "", err
	}

	var isNew = false
	if user == nil {
		user = &model.User{
			UserId:     userId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			SourceCode: sourceCode,
		}
		err := repo.UserRepo.Save(user)
		if err != nil {
			return false, "", err
		}
		isNew = true
	}

	resp, err := rpc.GetLogicIntClient().GetDevice(ctx, &pb.GetDeviceReq{DeviceId: deviceId})
	if err != nil {
		return false, "", err
	}
	token := util.RandString(40)
	// 业务 user_id
	err = repo.AuthRepo.Set(user.UserId, resp.Device.DeviceId, model.Device{
		Type:   resp.Device.Type,
		Token:  token,
		Expire: time.Now().AddDate(0, 3, 0).Unix(), // 默认有效期 3 个月
	})
	if err != nil {
		return false, "", err
	}

	return isNew, token, nil
}

// Auth 验证用户是否登录
func (*authService) Auth(ctx context.Context, userId string, deviceId int64, token string) error {
	device, err := repo.AuthRepo.Get(userId, deviceId)
	if err != nil {
		return err
	}

	if device == nil {
		return gerrors.ErrUnauthorized
	}

	if device.Expire < time.Now().Unix() {
		return gerrors.ErrUnauthorized
	}

	if device.Token != token {
		return gerrors.ErrUnauthorized
	}
	return nil
}
