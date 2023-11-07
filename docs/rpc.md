# 提供的的 RPC 功能方法

* 注册设备,用户需要绑定到设备上进行登录的 LogicExt.RegisterDevice
* 获取设备信息 LogicInt.GetDevice
* 用户登录设备离线 LogicInt.Offline
* 手机号验证码登录 BusinessExt.SignIn
* 获取登录用户信息 BusinessExt.GetUser
* 更新登录的用户信息 BusinessExt.UpdateUser
* 根据指定的key查询用户信息 BusinessExt.SearchUser
* 添加好友 LogicExt.AddFriend
* 同意添加好友请求 LogicExt.AgreeAddFriend
* 设置好友信息 LogicExt.SetFriend
* 获取好友列表 LogicExt.GetFriends
* 给好友发送消息 LogicExt.SendMessageToFriend
* 创建群组 LogicExt.CreateGroup
* 更新群组 LogicExt.UpdateGroup
* 获取指定群组信息 LogicExt.GetGroup
* 获取用户加入的所有群组 LogicExt.GetGroups
* 添加群组成员 LogicExt.AddGroupMembers
* 更新群组成员信息 LogicExt.UpdateGroupMember
* 移除群组成员 LogicExt.DeleteGroupMember
* 获取群组成员 LogicExt.GetGroupMembers
* 发送群组消息 LogicExt.SendMessageToGroup
* 长链接登录 LogicInt.ConnSignIn [这个是内部调用的，外部不允许使用]
* 房间（不需要创建的，临时保存到 redis里面的）
* 订阅房间
* 推送消息到房间

# 提供的事件回调

* 添加好友请求 code=110
* 同意添加好友 code=111
* 用户消息 code=100
* 群组消息 code=101
* 更新群组 code=120
* 添加群组成员 code=121
* 移除群组成员 code=122