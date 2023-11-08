package api

import (
	"context"

	app2 "gim/internal/business/domain/user/app"
	"gim/pkg/grpclib"
	"gim/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type BusinessExtServer struct {
	pb.UnsafeBusinessExtServer
}

// SignIn 根据 user_id 进行登录
func (s *BusinessExtServer) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	isNew, token, err := app2.AuthApp.SignIn(ctx, req.UserId, req.DeviceId, req.SourceCode)
	if err != nil {
		return nil, err
	}
	return &pb.SignInResp{
		IsNew: isNew,
		Token: token,
	}, nil
}

// GetUser 获取指定用户的信息
func (s *BusinessExtServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	user, err := app2.UserApp.Get(ctx, userId)
	return &pb.GetUserResp{User: user}, err
}

// UpdateUser 更新用户信息
func (s *BusinessExtServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}
	err = app2.UserApp.Update(ctx, userId, req)
	return new(emptypb.Empty), err
}

// SearchUser 根据关键词：手机号或者昵称 查询用户列表
func (s *BusinessExtServer) SearchUser(ctx context.Context, req *pb.SearchUserReq) (*pb.SearchUserResp, error) {
	users, err := app2.UserApp.Search(ctx, req.Key)
	return &pb.SearchUserResp{Users: users}, err
}
