package group

import (
	"context"

	"gim/internal/logic/domain/group/entity"
	"gim/internal/logic/domain/group/repo"
	"gim/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

// CreateGroup 创建群组
func (*app) CreateGroup(ctx context.Context, userId string, in *pb.CreateGroupReq) (string, error) {
	group := entity.CreateGroup(userId, in)
	err := repo.GroupRepo.Save(group)
	if err != nil {
		return "", err
	}
	return group.GroupId, nil
}

// GetGroup 获取群组信息
func (*app) GetGroup(ctx context.Context, groupId string) (*pb.Group, error) {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return nil, err
	}

	return group.ToProto(), nil
}

// GetUserGroups 获取用户加入的群组列表
func (*app) GetUserGroups(ctx context.Context, userId string) ([]*pb.Group, error) {
	groups, err := repo.GroupUserRepo.ListByUserId(userId)
	if err != nil {
		return nil, err
	}

	pbGroups := make([]*pb.Group, len(groups))
	for i := range groups {
		pbGroups[i] = groups[i].ToProto()
	}
	return pbGroups, nil
}

// Update 更新群组
func (*app) Update(ctx context.Context, userId string, update *pb.UpdateGroupReq) error {
	group, err := repo.GroupRepo.Get(update.GroupId)
	if err != nil {
		return err
	}

	err = group.Update(ctx, update)
	if err != nil {
		return err
	}

	err = repo.GroupRepo.Save(group)
	if err != nil {
		return err
	}
	// 删除对应的缓存信息
	_ = repo.GroupCache.Del(update.GroupId)

	// 推送更新群基础的事件
	err = group.PushUpdate(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

// AddMembers 添加群组成员
// 返回本地添加成功的用户 id 列表
func (*app) AddMembers(ctx context.Context, userId, groupId string, userIds []string) ([]string, error) {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return nil, err
	}
	_, addedIds, err := group.AddMembers(ctx, userIds)
	if err != nil {
		return nil, err
	}
	err = repo.GroupRepo.Save(group)
	if err != nil {
		return nil, err
	}

	err = group.PushAddMember(ctx, groupId, userId, addedIds)
	if err != nil {
		return nil, err
	}
	return addedIds, nil
}

// UpdateMember 更新群组用户
func (*app) UpdateMember(ctx context.Context, in *pb.UpdateGroupMemberReq) error {
	group, err := repo.GroupRepo.Get(in.GroupId)
	if err != nil {
		return err
	}
	err = group.UpdateMember(ctx, in)
	if err != nil {
		return err
	}
	err = repo.GroupRepo.Save(group)
	if err != nil {
		return err
	}
	return nil
}

// DeleteMember 删除群组成员
func (*app) DeleteMember(ctx context.Context, groupId, userId string, optId string) error {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return err
	}
	err = group.DeleteMember(ctx, userId)
	if err != nil {
		return err
	}
	err = repo.GroupRepo.Save(group)
	if err != nil {
		return err
	}

	err = group.PushDeleteMember(ctx, groupId, optId, userId)
	if err != nil {
		return err
	}
	return nil
}

// GetMembers 获取群组成员
func (*app) GetMembers(ctx context.Context, groupId string) ([]*pb.GroupMember, error) {
	group, err := repo.GroupRepo.Get(groupId)
	if err != nil {
		return nil, err
	}
	return group.GetMembers(ctx)
}

// SendMessage 发送群组消息
func (*app) SendMessage(ctx context.Context, fromDeviceID int64, fromUserID string, req *pb.SendMessageReq) (int64, error) {
	group, err := repo.GroupRepo.Get(req.ReceiverId)
	if err != nil {
		return 0, err
	}

	return group.SendMessage(ctx, fromDeviceID, fromUserID, req)
}
