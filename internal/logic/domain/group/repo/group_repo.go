package repo

import (
	"gim/internal/logic/domain/group/entity"
)

type groupRepo struct{}

var GroupRepo = new(groupRepo)

// Get 获取群组信息, 包含群成员信息
func (*groupRepo) Get(groupId string) (*entity.Group, error) {
	group, err := GroupCache.Get(groupId)
	if err != nil {
		return nil, err
	}
	if group != nil {
		return group, nil
	}

	group, err = GroupDao.Get(groupId)
	if err != nil {
		return nil, err
	}
	members, err := GroupUserRepo.ListUser(groupId)
	if err != nil {
		return nil, err
	}
	group.Members = members

	err = GroupCache.Set(group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// Save 保存群组信息
// 1. 群基础信息 2. 群成员信息
func (*groupRepo) Save(group *entity.Group) error {
	groupId := group.GroupId
	var err error
	members := group.Members
	for i := range members {
		members[i].GroupId = group.GroupId
		if members[i].UpdateType == entity.UpdateTypeUpdate {
			err = GroupUserRepo.Save(&(members[i]))
			if err != nil {
				return err
			}
			// TODO 这里的逻辑有问题，把新用户添加到群组中，会把旧的用户在添加一遍
			group.UserNum += 1
		}
		if members[i].UpdateType == entity.UpdateTypeDelete {
			err = GroupUserRepo.Delete(group.GroupId, members[i].UserId)
			if err != nil {
				return err
			}
			group.UserNum -= 1
		}
	}

	if groupId != "" {
		err = GroupCache.Del(groupId)
		if err != nil {
			return err
		}
	}
	// 保存更新后的群组信息
	err = GroupDao.Save(group)
	if err != nil {
		return err
	}
	return nil
}
