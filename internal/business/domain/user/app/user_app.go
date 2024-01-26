package app

import (
	"context"
	"time"

	"gim/internal/business/domain/user/repo"
	"gim/pkg/protocol/pb"
)

type userApp struct{}

var UserApp = new(userApp)

func (*userApp) Get(ctx context.Context, userId string) (*pb.User, error) {
	user, err := repo.UserRepo.Get(userId)
	return user.ToProto(), err
}

func (*userApp) Update(ctx context.Context, userId string, req *pb.UpdateUserReq) error {
	u, err := repo.UserRepo.Get(userId)
	if err != nil {
		return err
	}
	if u == nil {
		return nil
	}
	if req.Nickname != "" {
		u.Nickname = req.Nickname
	}
	if req.Sex != 0 {
		u.Sex = req.Sex
	}
	if req.AvatarUrl != "" {
		u.AvatarUrl = req.AvatarUrl
	}
	if req.Extra != "" {
		u.Extra = req.Extra
	}
	if req.Phone != "" {
		u.Phone = req.Phone
	}
	u.UpdateTime = time.Now()

	err = repo.UserRepo.Save(u)
	if err != nil {
		return err
	}
	return nil
}

// GetByIds 根据 id 批量获取用户信息
func (*userApp) GetByIds(ctx context.Context, userIds []string) (map[string]*pb.User, error) {
	users, err := repo.UserRepo.GetByIds(userIds)
	if err != nil {
		return nil, err
	}

	pbUsers := make(map[string]*pb.User, len(users))
	for i := range users {
		pbUsers[users[i].UserId] = users[i].ToProto()
	}
	return pbUsers, nil
}

func (*userApp) Search(ctx context.Context, key string) ([]*pb.User, error) {
	users, err := repo.UserRepo.Search(key)
	if err != nil {
		return nil, err
	}

	pbUsers := make([]*pb.User, len(users))
	for i, v := range users {
		pbUsers[i] = v.ToProto()
	}
	return pbUsers, nil
}
