package model

import (
	"time"

	"gim/pkg/protocol/pb"
)

// User 账户
type User struct {
	Id         int64     // 用户id
	UserId     string    // 用户业务唯一id
	Nickname   string    // 昵称
	Sex        int32     // 性别，0未知；1:男；2:女
	AvatarUrl  string    // 用户头像
	Extra      string    // 附加属性
	SourceCode string    // 来源编码
	Phone      string    // 手机号
	CreateTime time.Time // 创建时间
	UpdateTime time.Time // 更新时间
}

func (u *User) ToProto() *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{
		Id:         u.Id,
		UserId:     u.UserId,
		Nickname:   u.Nickname,
		Sex:        u.Sex,
		AvatarUrl:  u.AvatarUrl,
		Extra:      u.Extra,
		Phone:      u.Phone,
		CreateTime: u.CreateTime.Unix(),
		UpdateTime: u.UpdateTime.Unix(),
	}
}
