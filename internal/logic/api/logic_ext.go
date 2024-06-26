package api

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"gim/internal/logic/domain/device"
	"gim/internal/logic/domain/friend"
	"gim/internal/logic/domain/group"
	"gim/internal/logic/domain/room"
	"gim/pkg/grpclib"
	"gim/pkg/logger"
	"gim/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type LogicExtServer struct {
	pb.UnsafeLogicExtServer
}

// RegisterDevice 注册设备
func (*LogicExtServer) RegisterDevice(ctx context.Context, in *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	deviceId, err := device.App.Register(ctx, in)
	return &pb.RegisterDeviceResp{DeviceId: deviceId}, err
}

func (*LogicExtServer) GetDeviceById(ctx context.Context, in *pb.GetDeviceByIdReq) (*pb.GetDeviceByIdResp, error) {
	deviceInfo, err := device.App.GetDevice(ctx, in.DeviceId)
	return &pb.GetDeviceByIdResp{
		DeviceId: deviceInfo.DeviceId,
		Type:     deviceInfo.Type,
		Brand:    deviceInfo.Brand,
	}, err
}

// PushRoom  推送房间
func (s *LogicExtServer) PushRoom(ctx context.Context, req *pb.PushRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, room.App.Push(ctx, req)
}

// SendMessageToFriend 发送好友消息
// 返回的序列为：userID+seq
func (*LogicExtServer) SendMessageToFriend(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	logger.Logger.Debug("start send message:", zap.String("time", time.Now().String()), zap.Int64("send_time", in.SendTime))
	userId, deviceId, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	seq, err := friend.App.SendToFriend(ctx, deviceId, userId, in)
	if err != nil {
		return nil, err
	}
	logger.Logger.Debug("end send message:", zap.String("time", time.Now().String()), zap.Int64("send_time", in.SendTime))
	return &pb.SendMessageResp{
		Seq:     seq,
		UserSeq: fmt.Sprintf("%s_%d", userId, seq),
	}, nil
}

// AddFriend 添加好友
// 判断还有是否存在
// 修改为不需要同意，直接就通过了
func (s *LogicExtServer) AddFriend(ctx context.Context, in *pb.AddFriendReq) (*emptypb.Empty, error) {
	// userId, _, err := grpclib.GetCtxData(ctx)
	// if err != nil {
	//     return nil, err
	// }

	err := friend.App.AddFriend(ctx, in.UserId, in.FriendId, in.Remarks, in.Description)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// AgreeAddFriend  同意好友请求
func (s *LogicExtServer) AgreeAddFriend(ctx context.Context, in *pb.AgreeAddFriendReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = friend.App.AgreeAddFriend(ctx, userId, in.UserId, in.Remarks)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *LogicExtServer) SetFriend(ctx context.Context, req *pb.SetFriendReq) (*pb.SetFriendResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = friend.App.SetFriend(ctx, userId, req)
	if err != nil {
		return nil, err
	}
	return &pb.SetFriendResp{}, nil
}

func (s *LogicExtServer) GetFriends(ctx context.Context, in *pb.GetFriendsReq) (*pb.GetFriendsResp, error) {
	// userId, _, err := grpclib.GetCtxData(ctx)
	// if err != nil {
	//     return nil, err
	// }
	friends, err := friend.App.List(ctx, in.UserId)
	return &pb.GetFriendsResp{Friends: friends}, err
}

// SendMessageToGroup 发送群组消息
func (*LogicExtServer) SendMessageToGroup(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	userId, deviceId, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	seq, err := group.App.SendMessage(ctx, deviceId, userId, in)
	if err != nil {
		return nil, err
	}
	return &pb.SendMessageResp{
		UserSeq: fmt.Sprintf("%s_%d", userId, seq),
		Seq:     seq,
	}, nil
}

// CreateGroup 创建群组
// 传的参数 member_ids 是不包含创建人的，创建人默认入群 而且是管理员
// 注意，这里不判断 添加到群成员的是否真是存在
// 群相关信息添加到缓存
func (*LogicExtServer) CreateGroup(ctx context.Context, in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}
	groupId, err := group.App.CreateGroup(ctx, userId, in)
	return &pb.CreateGroupResp{GroupId: groupId}, err
}

// UpdateGroup 更新群组，这里只更新基础的信息
func (*LogicExtServer) UpdateGroup(ctx context.Context, in *pb.UpdateGroupReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, group.App.Update(ctx, userId, in)
}

// GetGroup 获取群组信息
func (*LogicExtServer) GetGroup(ctx context.Context, in *pb.GetGroupReq) (*pb.GetGroupResp, error) {
	group, err := group.App.GetGroup(ctx, in.GroupId)
	return &pb.GetGroupResp{Group: group}, err
}

// GetGroups 获取当前用户加入的所有群组
// 不包含群成员信息
func (*LogicExtServer) GetGroups(ctx context.Context, in *emptypb.Empty) (*pb.GetGroupsResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	groups, err := group.App.GetUserGroups(ctx, userId)
	return &pb.GetGroupsResp{Groups: groups}, err
}

// AddGroupMembers 添加群成员到指定群中
func (s *LogicExtServer) AddGroupMembers(ctx context.Context, in *pb.AddGroupMembersReq) (*pb.AddGroupMembersResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	userIds, err := group.App.AddMembers(ctx, userId, in.GroupId, in.UserIds)
	return &pb.AddGroupMembersResp{UserIds: userIds}, err
}

// UpdateGroupMember 更新群组成员信息
// 针对的是已经在群组中的用户 进行更更新：成员类型，备注，附加字段 这些信息
func (*LogicExtServer) UpdateGroupMember(ctx context.Context, in *pb.UpdateGroupMemberReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, group.App.UpdateMember(ctx, in)
}

// DeleteGroupMember 删除群组成员
// 这里只能单个成员进行删除
func (*LogicExtServer) DeleteGroupMember(ctx context.Context, in *pb.DeleteGroupMemberReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = group.App.DeleteMember(ctx, in.GroupId, in.UserId, userId)
	return &emptypb.Empty{}, err
}

// GetGroupMembers 获取群组成员信息
// 不分页，直接获取的
func (s *LogicExtServer) GetGroupMembers(ctx context.Context, in *pb.GetGroupMembersReq) (*pb.GetGroupMembersResp, error) {
	members, err := group.App.GetMembers(ctx, in.GroupId)
	return &pb.GetGroupMembersResp{Members: members}, err
}
