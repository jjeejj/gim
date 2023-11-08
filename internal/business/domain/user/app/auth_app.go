package app

import (
	"context"

	"gim/internal/business/domain/user/service"
)

type authApp struct{}

var AuthApp = new(authApp)

// SignIn 手机验证码+设备登录登录
func (*authApp) SignIn(ctx context.Context, userId string, deviceId int64, sourceCode string) (bool, string, error) {
	return service.AuthService.SignIn(ctx, userId, deviceId, sourceCode)
}

// Auth 验证用户是否登录
func (*authApp) Auth(ctx context.Context, userId string, deviceId int64, token string) error {
	return service.AuthService.Auth(ctx, userId, deviceId, token)
}
