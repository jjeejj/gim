package friend

import (
	"context"
	"time"

	"gim/pkg/protocol/pb"
)

type app struct{}

var App = new(app)

// List 获取好友列表
func (s *app) List(ctx context.Context, userId string) ([]*pb.Friend, error) {
	return Service.List(ctx, userId)
}

// AddFriend 添加好友
func (*app) AddFriend(ctx context.Context, userId, friendId string, remarks, description string) error {
	return Service.AddFriend(ctx, userId, friendId, remarks, description)
}

// AgreeAddFriend 同意添加好友
func (*app) AgreeAddFriend(ctx context.Context, userId, friendId string, remarks string) error {
	return Service.AgreeAddFriend(ctx, userId, friendId, remarks)
}

// SetFriend 设置好友信息
func (*app) SetFriend(ctx context.Context, userId string, req *pb.SetFriendReq) error {
	friend, err := Repo.Get(userId, req.FriendId)
	if err != nil {
		return err
	}
	if friend == nil {
		return nil
	}

	friend.Remarks = req.Remarks
	friend.Extra = req.Extra
	friend.UpdateTime = time.Now()

	err = Repo.Save(friend)
	if err != nil {
		return err
	}
	return nil
}

// SendToFriend 消息发送至好友
// 不添加还有也可以发送消息
func (*app) SendToFriend(ctx context.Context, fromDeviceID int64, fromUserID string, req *pb.SendMessageReq) (int64, error) {
	return Service.SendToFriend(ctx, fromDeviceID, fromUserID, req)
}
