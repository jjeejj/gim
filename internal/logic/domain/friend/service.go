package friend

import (
	"context"
	"time"

	"gim/internal/logic/proxy"
	"gim/pkg/gerrors"
	"gim/pkg/protocol/pb"
	"gim/pkg/rpc"

	"google.golang.org/protobuf/proto"
)

type service struct{}

var Service = new(service)

// List 获取好友列表
func (s *service) List(ctx context.Context, userId string) ([]*pb.Friend, error) {
	friends, err := Repo.List(userId, FriendStatusAgree)
	if err != nil {
		return nil, err
	}

	userIds := make(map[string]int32, len(friends))
	for i := range friends {
		userIds[friends[i].FriendId] = 0
	}
	resp, err := rpc.GetBusinessIntClient().GetUsers(ctx, &pb.GetUsersReq{UserIds: userIds})
	if err != nil {
		return nil, err
	}

	var infos = make([]*pb.Friend, len(friends))
	for i := range friends {
		friend := pb.Friend{
			UserId:  friends[i].FriendId,
			Remarks: friends[i].Remarks,
			Extra:   friends[i].Extra,
		}

		user, ok := resp.Users[friends[i].FriendId]
		if ok {
			friend.Nickname = user.Nickname
			friend.Sex = user.Sex
			friend.AvatarUrl = user.AvatarUrl
			friend.UserExtra = user.Extra
		}
		infos[i] = &friend
	}

	return infos, nil
}

// AddFriend 添加好友
func (*service) AddFriend(ctx context.Context, userId, friendId string, remarks, description string) error {
	// 判断添加的好友是否存在
	friendUserResp, err := rpc.GetBusinessIntClient().GetUser(ctx, &pb.GetUserReq{UserId: friendId})
	if err != nil {
		return err
	}
	if friendUserResp == nil || friendUserResp.User == nil {
		return gerrors.ErrUserNotFound
	}
	// 判断是否已经是好友关系
	friend, err := Repo.Get(userId, friendId)
	if err != nil {
		return err
	}
	if friend != nil {
		if friend.Status == FriendStatusApply {
			return nil
		}
		if friend.Status == FriendStatusAgree {
			return gerrors.ErrAlreadyIsFriend
		}
	}

	now := time.Now()
	err = Repo.Save(&Friend{
		UserId:     userId,
		FriendId:   friendId,
		Remarks:    remarks,
		Status:     FriendStatusApply,
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		return err
	}

	resp, err := rpc.GetBusinessIntClient().GetUser(ctx, &pb.GetUserReq{UserId: userId})
	if err != nil {
		return err
	}

	_, err = proxy.PushToUser(ctx, friendId, pb.PushCode_PC_ADD_FRIEND, &pb.AddFriendPush{
		FriendId:    userId,
		Nickname:    resp.User.Nickname,
		AvatarUrl:   resp.User.AvatarUrl,
		Description: description,
	}, true)
	if err != nil {
		return err
	}
	return nil
}

// AgreeAddFriend 同意添加好友
func (*service) AgreeAddFriend(ctx context.Context, userId, friendId string, remarks string) error {
	friend, err := Repo.Get(friendId, userId)
	if err != nil {
		return err
	}
	if friend == nil {
		return gerrors.ErrBadRequest
	}
	if friend.Status == FriendStatusAgree {
		return nil
	}
	friend.Status = FriendStatusAgree
	err = Repo.Save(friend)
	if err != nil {
		return err
	}

	now := time.Now()
	err = Repo.Save(&Friend{
		UserId:     userId,
		FriendId:   friendId,
		Remarks:    remarks,
		Status:     FriendStatusAgree,
		CreateTime: now,
		UpdateTime: now,
	})
	if err != nil {
		return err
	}

	resp, err := rpc.GetBusinessIntClient().GetUser(ctx, &pb.GetUserReq{UserId: userId})
	if err != nil {
		return err
	}

	_, err = proxy.PushToUser(ctx, friendId, pb.PushCode_PC_AGREE_ADD_FRIEND, &pb.AgreeAddFriendPush{
		FriendId:  userId,
		Nickname:  resp.User.Nickname,
		AvatarUrl: resp.User.AvatarUrl,
	}, true)
	if err != nil {
		return err
	}
	return nil
}

// SendToFriend 消息发送至好友
func (*service) SendToFriend(ctx context.Context, fromDeviceID int64, fromUserID string, req *pb.SendMessageReq) (int64, error) {
	sender, err := rpc.GetSender(fromDeviceID, fromUserID)
	if err != nil {
		return 0, err
	}

	// 发给发送者
	push := pb.UserMessagePush{
		Sender:     sender,
		ReceiverId: req.ReceiverId,
		Content:    req.Content,
	}
	bytes, err := proto.Marshal(&push)
	if err != nil {
		return 0, err
	}

	msg := &pb.Message{
		Code:     int32(pb.PushCode_PC_USER_MESSAGE),
		Content:  bytes,
		SendTime: req.SendTime,
	}
	// 给发送着通知
	seq, err := proxy.MessageProxy.SendToUser(ctx, fromDeviceID, fromUserID, msg, true)
	if err != nil {
		return 0, err
	}

	// 发给接收者
	_, err = proxy.MessageProxy.SendToUser(ctx, fromDeviceID, req.ReceiverId, msg, true)
	if err != nil {
		return 0, err
	}

	return seq, nil
}
