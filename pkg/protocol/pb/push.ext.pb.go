// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.25.0
// source: push.ext.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PushCode int32

const (
	PushCode_PC_ADD_DEFAULT         PushCode = 0
	PushCode_PC_USER_MESSAGE        PushCode = 100 // 用户消息
	PushCode_PC_GROUP_MESSAGE       PushCode = 101 // 群组消息
	PushCode_PC_ADD_FRIEND          PushCode = 110 // 添加好友请求
	PushCode_PC_AGREE_ADD_FRIEND    PushCode = 111 // 同意添加好友
	PushCode_PC_UPDATE_GROUP        PushCode = 120 // 更新群组
	PushCode_PC_ADD_GROUP_MEMBERS   PushCode = 121 // 添加群组成员
	PushCode_PC_REMOVE_GROUP_MEMBER PushCode = 122 // 移除群组成员
)

// Enum value maps for PushCode.
var (
	PushCode_name = map[int32]string{
		0:   "PC_ADD_DEFAULT",
		100: "PC_USER_MESSAGE",
		101: "PC_GROUP_MESSAGE",
		110: "PC_ADD_FRIEND",
		111: "PC_AGREE_ADD_FRIEND",
		120: "PC_UPDATE_GROUP",
		121: "PC_ADD_GROUP_MEMBERS",
		122: "PC_REMOVE_GROUP_MEMBER",
	}
	PushCode_value = map[string]int32{
		"PC_ADD_DEFAULT":         0,
		"PC_USER_MESSAGE":        100,
		"PC_GROUP_MESSAGE":       101,
		"PC_ADD_FRIEND":          110,
		"PC_AGREE_ADD_FRIEND":    111,
		"PC_UPDATE_GROUP":        120,
		"PC_ADD_GROUP_MEMBERS":   121,
		"PC_REMOVE_GROUP_MEMBER": 122,
	}
)

func (x PushCode) Enum() *PushCode {
	p := new(PushCode)
	*p = x
	return p
}

func (x PushCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushCode) Descriptor() protoreflect.EnumDescriptor {
	return file_push_ext_proto_enumTypes[0].Descriptor()
}

func (PushCode) Type() protoreflect.EnumType {
	return &file_push_ext_proto_enumTypes[0]
}

func (x PushCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushCode.Descriptor instead.
func (PushCode) EnumDescriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{0}
}

type Sender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 发送者 业务id
	DeviceId  int64  `protobuf:"varint,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`   // 发送者设备id
	AvatarUrl string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 昵称
	Nickname  string `protobuf:"bytes,5,opt,name=nickname,proto3" json:"nickname,omitempty"`                    // 头像
	Extra     string `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`                          // 扩展字段
}

func (x *Sender) Reset() {
	*x = Sender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sender) ProtoMessage() {}

func (x *Sender) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sender.ProtoReflect.Descriptor instead.
func (*Sender) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{0}
}

func (x *Sender) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Sender) GetDeviceId() int64 {
	if x != nil {
		return x.DeviceId
	}
	return 0
}

func (x *Sender) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *Sender) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Sender) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

// 用户消息 MC_USER_MESSAGE = 100
type UserMessagePush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sender     *Sender `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	ReceiverId string  `protobuf:"bytes,2,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"` // 用户业务id或者群组id
	Content    []byte  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                         // 用户发送的消息内容
}

func (x *UserMessagePush) Reset() {
	*x = UserMessagePush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMessagePush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMessagePush) ProtoMessage() {}

func (x *UserMessagePush) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMessagePush.ProtoReflect.Descriptor instead.
func (*UserMessagePush) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{1}
}

func (x *UserMessagePush) GetSender() *Sender {
	if x != nil {
		return x.Sender
	}
	return nil
}

func (x *UserMessagePush) GetReceiverId() string {
	if x != nil {
		return x.ReceiverId
	}
	return ""
}

func (x *UserMessagePush) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

// 添加好友 PC_ADD_FRIEND = 110
type AddFriendPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId    string `protobuf:"bytes,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`    // 好友业务id
	Nickname    string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                    // 昵称
	AvatarUrl   string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 头像
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`              // 描述
}

func (x *AddFriendPush) Reset() {
	*x = AddFriendPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFriendPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFriendPush) ProtoMessage() {}

func (x *AddFriendPush) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFriendPush.ProtoReflect.Descriptor instead.
func (*AddFriendPush) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{2}
}

func (x *AddFriendPush) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

func (x *AddFriendPush) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AddFriendPush) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *AddFriendPush) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// 同意 添加好友 PC_AGREE_ADD_FRIEND = 111
type AgreeAddFriendPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FriendId  string `protobuf:"bytes,1,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`    // 好友id
	Nickname  string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                    // 昵称
	AvatarUrl string `protobuf:"bytes,3,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 头像
}

func (x *AgreeAddFriendPush) Reset() {
	*x = AgreeAddFriendPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgreeAddFriendPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgreeAddFriendPush) ProtoMessage() {}

func (x *AgreeAddFriendPush) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgreeAddFriendPush.ProtoReflect.Descriptor instead.
func (*AgreeAddFriendPush) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{3}
}

func (x *AgreeAddFriendPush) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

func (x *AgreeAddFriendPush) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *AgreeAddFriendPush) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

// 更新群组 PC_UPDATE_GROUP = 120
type UpdateGroupPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptId        string `protobuf:"bytes,1,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`             // 操作人用户id
	OptName      string `protobuf:"bytes,2,opt,name=opt_name,json=optName,proto3" json:"opt_name,omitempty"`       // 操作人昵称
	Name         string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                            // 群组名称
	AvatarUrl    string `protobuf:"bytes,4,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"` // 群组头像
	Introduction string `protobuf:"bytes,5,opt,name=introduction,proto3" json:"introduction,omitempty"`            // 群组简介
	Extra        string `protobuf:"bytes,6,opt,name=extra,proto3" json:"extra,omitempty"`                          // 附加字段
	GroupId      string `protobuf:"bytes,7,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`       // 变更的群组id
}

func (x *UpdateGroupPush) Reset() {
	*x = UpdateGroupPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupPush) ProtoMessage() {}

func (x *UpdateGroupPush) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupPush.ProtoReflect.Descriptor instead.
func (*UpdateGroupPush) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateGroupPush) GetOptId() string {
	if x != nil {
		return x.OptId
	}
	return ""
}

func (x *UpdateGroupPush) GetOptName() string {
	if x != nil {
		return x.OptName
	}
	return ""
}

func (x *UpdateGroupPush) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateGroupPush) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *UpdateGroupPush) GetIntroduction() string {
	if x != nil {
		return x.Introduction
	}
	return ""
}

func (x *UpdateGroupPush) GetExtra() string {
	if x != nil {
		return x.Extra
	}
	return ""
}

func (x *UpdateGroupPush) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// 添加群组成员 PC_AGREE_ADD_GROUPS = 121
type AddGroupMembersPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptId   string         `protobuf:"bytes,1,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`       // 操作人用户id
	OptName string         `protobuf:"bytes,2,opt,name=opt_name,json=optName,proto3" json:"opt_name,omitempty"` // 操作人昵称
	Members []*GroupMember `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`                // 群组成员
	GroupId string         `protobuf:"bytes,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"` // 变更的群组id
}

func (x *AddGroupMembersPush) Reset() {
	*x = AddGroupMembersPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddGroupMembersPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddGroupMembersPush) ProtoMessage() {}

func (x *AddGroupMembersPush) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddGroupMembersPush.ProtoReflect.Descriptor instead.
func (*AddGroupMembersPush) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{5}
}

func (x *AddGroupMembersPush) GetOptId() string {
	if x != nil {
		return x.OptId
	}
	return ""
}

func (x *AddGroupMembersPush) GetOptName() string {
	if x != nil {
		return x.OptName
	}
	return ""
}

func (x *AddGroupMembersPush) GetMembers() []*GroupMember {
	if x != nil {
		return x.Members
	}
	return nil
}

func (x *AddGroupMembersPush) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

// 删除群组成员 PC_REMOVE_GROUP_MEMBER = 122
type RemoveGroupMemberPush struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptId         string `protobuf:"bytes,1,opt,name=opt_id,json=optId,proto3" json:"opt_id,omitempty"`                           // 操作人用户业务id
	OptName       string `protobuf:"bytes,2,opt,name=opt_name,json=optName,proto3" json:"opt_name,omitempty"`                     // 操作人昵称
	DeletedUserId string `protobuf:"bytes,3,opt,name=deleted_user_id,json=deletedUserId,proto3" json:"deleted_user_id,omitempty"` // 被删除的成员id
	GroupId       string `protobuf:"bytes,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`                     // 变更的群组id
}

func (x *RemoveGroupMemberPush) Reset() {
	*x = RemoveGroupMemberPush{}
	if protoimpl.UnsafeEnabled {
		mi := &file_push_ext_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveGroupMemberPush) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveGroupMemberPush) ProtoMessage() {}

func (x *RemoveGroupMemberPush) ProtoReflect() protoreflect.Message {
	mi := &file_push_ext_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveGroupMemberPush.ProtoReflect.Descriptor instead.
func (*RemoveGroupMemberPush) Descriptor() ([]byte, []int) {
	return file_push_ext_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveGroupMemberPush) GetOptId() string {
	if x != nil {
		return x.OptId
	}
	return ""
}

func (x *RemoveGroupMemberPush) GetOptName() string {
	if x != nil {
		return x.OptName
	}
	return ""
}

func (x *RemoveGroupMemberPush) GetDeletedUserId() string {
	if x != nil {
		return x.DeletedUserId
	}
	return ""
}

func (x *RemoveGroupMemberPush) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

var File_push_ext_proto protoreflect.FileDescriptor

var file_push_ext_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x75, 0x73, 0x68, 0x2e, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2e, 0x65, 0x78, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x22, 0x70, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x75, 0x73, 0x68, 0x12, 0x22, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x12, 0x41, 0x67, 0x72, 0x65, 0x65, 0x41, 0x64,
	0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x50, 0x75, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x55, 0x72, 0x6c, 0x22, 0xcb, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x70, 0x74, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x70, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x70, 0x74, 0x49, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x22, 0x8c, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x50, 0x75, 0x73, 0x68, 0x12, 0x15, 0x0a, 0x06, 0x6f,
	0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x70, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a,
	0x0f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x2a, 0xc0, 0x01, 0x0a, 0x08, 0x50, 0x75, 0x73, 0x68, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x0e, 0x50, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10,
	0x00, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x43, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x53,
	0x53, 0x41, 0x47, 0x45, 0x10, 0x64, 0x12, 0x14, 0x0a, 0x10, 0x50, 0x43, 0x5f, 0x47, 0x52, 0x4f,
	0x55, 0x50, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x50, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x6e, 0x12,
	0x17, 0x0a, 0x13, 0x50, 0x43, 0x5f, 0x41, 0x47, 0x52, 0x45, 0x45, 0x5f, 0x41, 0x44, 0x44, 0x5f,
	0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x6f, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x43, 0x5f, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x78, 0x12, 0x18, 0x0a,
	0x14, 0x50, 0x43, 0x5f, 0x41, 0x44, 0x44, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45,
	0x4d, 0x42, 0x45, 0x52, 0x53, 0x10, 0x79, 0x12, 0x1a, 0x0a, 0x16, 0x50, 0x43, 0x5f, 0x52, 0x45,
	0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45,
	0x52, 0x10, 0x7a, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x69, 0x6d, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_push_ext_proto_rawDescOnce sync.Once
	file_push_ext_proto_rawDescData = file_push_ext_proto_rawDesc
)

func file_push_ext_proto_rawDescGZIP() []byte {
	file_push_ext_proto_rawDescOnce.Do(func() {
		file_push_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_push_ext_proto_rawDescData)
	})
	return file_push_ext_proto_rawDescData
}

var file_push_ext_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_push_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_push_ext_proto_goTypes = []interface{}{
	(PushCode)(0),                 // 0: pb.PushCode
	(*Sender)(nil),                // 1: pb.Sender
	(*UserMessagePush)(nil),       // 2: pb.UserMessagePush
	(*AddFriendPush)(nil),         // 3: pb.AddFriendPush
	(*AgreeAddFriendPush)(nil),    // 4: pb.AgreeAddFriendPush
	(*UpdateGroupPush)(nil),       // 5: pb.UpdateGroupPush
	(*AddGroupMembersPush)(nil),   // 6: pb.AddGroupMembersPush
	(*RemoveGroupMemberPush)(nil), // 7: pb.RemoveGroupMemberPush
	(*GroupMember)(nil),           // 8: pb.GroupMember
}
var file_push_ext_proto_depIdxs = []int32{
	1, // 0: pb.UserMessagePush.sender:type_name -> pb.Sender
	8, // 1: pb.AddGroupMembersPush.members:type_name -> pb.GroupMember
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_push_ext_proto_init() }
func file_push_ext_proto_init() {
	if File_push_ext_proto != nil {
		return
	}
	file_logic_ext_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_push_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sender); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_ext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMessagePush); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_ext_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFriendPush); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_ext_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgreeAddFriendPush); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_ext_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupPush); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_ext_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddGroupMembersPush); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_push_ext_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveGroupMemberPush); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_push_ext_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_push_ext_proto_goTypes,
		DependencyIndexes: file_push_ext_proto_depIdxs,
		EnumInfos:         file_push_ext_proto_enumTypes,
		MessageInfos:      file_push_ext_proto_msgTypes,
	}.Build()
	File_push_ext_proto = out.File
	file_push_ext_proto_rawDesc = nil
	file_push_ext_proto_goTypes = nil
	file_push_ext_proto_depIdxs = nil
}
