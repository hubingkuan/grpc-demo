// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: friend/friend.proto

package friend

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 加好友的途径
type AddFriendType int32

const (
	AddFriendType_NormalAdd AddFriendType = 0 //普通加好友
	AddFriendType_LetterAdd AddFriendType = 1 //留言板加好友
)

// Enum value maps for AddFriendType.
var (
	AddFriendType_name = map[int32]string{
		0: "NormalAdd",
		1: "LetterAdd",
	}
	AddFriendType_value = map[string]int32{
		"NormalAdd": 0,
		"LetterAdd": 1,
	}
)

func (x AddFriendType) Enum() *AddFriendType {
	p := new(AddFriendType)
	*p = x
	return p
}

func (x AddFriendType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AddFriendType) Descriptor() protoreflect.EnumDescriptor {
	return file_friend_friend_proto_enumTypes[0].Descriptor()
}

func (AddFriendType) Type() protoreflect.EnumType {
	return &file_friend_friend_proto_enumTypes[0]
}

func (x AddFriendType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AddFriendType.Descriptor instead.
func (AddFriendType) EnumDescriptor() ([]byte, []int) {
	return file_friend_friend_proto_rawDescGZIP(), []int{0}
}

type DegreeMethod int32

const (
	DegreeMethod_Other          DegreeMethod = 0
	DegreeMethod_PrivateChat    DegreeMethod = 1
	DegreeMethod_GiftGiveMethod DegreeMethod = 2
)

// Enum value maps for DegreeMethod.
var (
	DegreeMethod_name = map[int32]string{
		0: "Other",
		1: "PrivateChat",
		2: "GiftGiveMethod",
	}
	DegreeMethod_value = map[string]int32{
		"Other":          0,
		"PrivateChat":    1,
		"GiftGiveMethod": 2,
	}
)

func (x DegreeMethod) Enum() *DegreeMethod {
	p := new(DegreeMethod)
	*p = x
	return p
}

func (x DegreeMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DegreeMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_friend_friend_proto_enumTypes[1].Descriptor()
}

func (DegreeMethod) Type() protoreflect.EnumType {
	return &file_friend_friend_proto_enumTypes[1]
}

func (x DegreeMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DegreeMethod.Descriptor instead.
func (DegreeMethod) EnumDescriptor() ([]byte, []int) {
	return file_friend_friend_proto_rawDescGZIP(), []int{1}
}

// 好友基本信息
type FriendBaseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId     uint64        `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`                          // 玩家ID
	Name         string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                   // 昵称
	IsOnline     bool          `protobuf:"varint,3,opt,name=isOnline,proto3" json:"isOnline,omitempty"`                          // 是否在线
	Frame        string        `protobuf:"bytes,4,opt,name=frame,proto3" json:"frame,omitempty"`                                 // 相框 前缀必须是foo
	Head         uint32        `protobuf:"varint,5,opt,name=head,proto3" json:"head,omitempty"`                                  // 头像 必须在指定值1,2,3,4,5内
	Model        uint32        `protobuf:"varint,6,opt,name=model,proto3" json:"model,omitempty"`                                // 模型
	Tag          string        `protobuf:"bytes,7,opt,name=tag,proto3" json:"tag,omitempty"`                                     // 备注 最大长度1 最小长度10
	Offline      int64         `protobuf:"varint,8,opt,name=offline,proto3" json:"offline,omitempty"`                            // 离线时间
	FriendDegree int32         `protobuf:"varint,9,opt,name=friendDegree,proto3" json:"friendDegree,omitempty"`                  // 好友度
	AddType      AddFriendType `protobuf:"varint,10,opt,name=addType,proto3,enum=friend.AddFriendType" json:"addType,omitempty"` // 加好友的途径 字段必须是指定的枚举值
	BaseLevel    int32         `protobuf:"varint,11,opt,name=baseLevel,proto3" json:"baseLevel,omitempty"`                       // 角色等级
	Email        string        `protobuf:"bytes,12,opt,name=email,proto3" json:"email,omitempty"`
	X            bool          `protobuf:"varint,13,opt,name=x,proto3" json:"x,omitempty"`           // 必须是指定的值 true
	Xx           []float32     `protobuf:"fixed32,14,rep,packed,name=xx,proto3" json:"xx,omitempty"` // 重复的值必须是唯一的
}

func (x *FriendBaseInfo) Reset() {
	*x = FriendBaseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_friend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendBaseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendBaseInfo) ProtoMessage() {}

func (x *FriendBaseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_friend_friend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendBaseInfo.ProtoReflect.Descriptor instead.
func (*FriendBaseInfo) Descriptor() ([]byte, []int) {
	return file_friend_friend_proto_rawDescGZIP(), []int{0}
}

func (x *FriendBaseInfo) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *FriendBaseInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FriendBaseInfo) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *FriendBaseInfo) GetFrame() string {
	if x != nil {
		return x.Frame
	}
	return ""
}

func (x *FriendBaseInfo) GetHead() uint32 {
	if x != nil {
		return x.Head
	}
	return 0
}

func (x *FriendBaseInfo) GetModel() uint32 {
	if x != nil {
		return x.Model
	}
	return 0
}

func (x *FriendBaseInfo) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *FriendBaseInfo) GetOffline() int64 {
	if x != nil {
		return x.Offline
	}
	return 0
}

func (x *FriendBaseInfo) GetFriendDegree() int32 {
	if x != nil {
		return x.FriendDegree
	}
	return 0
}

func (x *FriendBaseInfo) GetAddType() AddFriendType {
	if x != nil {
		return x.AddType
	}
	return AddFriendType_NormalAdd
}

func (x *FriendBaseInfo) GetBaseLevel() int32 {
	if x != nil {
		return x.BaseLevel
	}
	return 0
}

func (x *FriendBaseInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FriendBaseInfo) GetX() bool {
	if x != nil {
		return x.X
	}
	return false
}

func (x *FriendBaseInfo) GetXx() []float32 {
	if x != nil {
		return x.Xx
	}
	return nil
}

type RadarSearchPlayerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Distance    float32 `protobuf:"fixed32,1,opt,name=distance,proto3" json:"distance,omitempty"`
	PlayerId    uint64  `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	BubbleFrame uint32  `protobuf:"varint,3,opt,name=bubbleFrame,proto3" json:"bubbleFrame,omitempty"` // 气泡框
	Head        uint32  `protobuf:"varint,4,opt,name=head,proto3" json:"head,omitempty"`               // 头像
	HeadFrame   uint32  `protobuf:"varint,5,opt,name=headFrame,proto3" json:"headFrame,omitempty"`     // 头像框
	NickName    string  `protobuf:"bytes,6,opt,name=nickName,proto3" json:"nickName,omitempty"`        // 昵称
}

func (x *RadarSearchPlayerInfo) Reset() {
	*x = RadarSearchPlayerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_friend_friend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RadarSearchPlayerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RadarSearchPlayerInfo) ProtoMessage() {}

func (x *RadarSearchPlayerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_friend_friend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RadarSearchPlayerInfo.ProtoReflect.Descriptor instead.
func (*RadarSearchPlayerInfo) Descriptor() ([]byte, []int) {
	return file_friend_friend_proto_rawDescGZIP(), []int{1}
}

func (x *RadarSearchPlayerInfo) GetDistance() float32 {
	if x != nil {
		return x.Distance
	}
	return 0
}

func (x *RadarSearchPlayerInfo) GetPlayerId() uint64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *RadarSearchPlayerInfo) GetBubbleFrame() uint32 {
	if x != nil {
		return x.BubbleFrame
	}
	return 0
}

func (x *RadarSearchPlayerInfo) GetHead() uint32 {
	if x != nil {
		return x.Head
	}
	return 0
}

func (x *RadarSearchPlayerInfo) GetHeadFrame() uint32 {
	if x != nil {
		return x.HeadFrame
	}
	return 0
}

func (x *RadarSearchPlayerInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

var File_friend_friend_proto protoreflect.FileDescriptor

var file_friend_friend_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x04, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x42,
	0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x32, 0x03,
	0x20, 0xe7, 0x07, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x42, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2e, 0xfa, 0x42, 0x2b,
	0x72, 0x29, 0x28, 0x80, 0x02, 0x32, 0x24, 0x5e, 0x5b, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x41,
	0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x28, 0x20, 0x5b, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x5d,
	0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x5d, 0x2b, 0x29, 0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a,
	0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x72, 0x05, 0x3a, 0x03, 0x66, 0x6f, 0x6f, 0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0f, 0xfa,
	0x42, 0x0c, 0x2a, 0x0a, 0x30, 0x01, 0x30, 0x02, 0x30, 0x03, 0x30, 0x04, 0x30, 0x05, 0x52, 0x04,
	0x68, 0x65, 0x61, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x2a, 0x04, 0x18, 0x5a, 0x28, 0x32, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x0a, 0x52, 0x03, 0x74,
	0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x44, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x12, 0x39, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x07, 0x61, 0x64, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x62, 0x61, 0x73, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60,
	0x01, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x15, 0x0a, 0x01, 0x78, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x6a, 0x02, 0x08, 0x01, 0x52, 0x01, 0x78, 0x12,
	0x18, 0x0a, 0x02, 0x78, 0x78, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x02, 0x42, 0x08, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x18, 0x01, 0x52, 0x02, 0x78, 0x78, 0x22, 0xbf, 0x01, 0x0a, 0x15, 0x52, 0x61,
	0x64, 0x61, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62,
	0x75, 0x62, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x62, 0x75, 0x62, 0x62, 0x6c, 0x65, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x65, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x68, 0x65, 0x61,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x68, 0x65, 0x61, 0x64, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x2a, 0x2d, 0x0a, 0x0d, 0x41,
	0x64, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09,
	0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x41, 0x64, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4c,
	0x65, 0x74, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x10, 0x01, 0x2a, 0x3e, 0x0a, 0x0c, 0x44, 0x65,
	0x67, 0x72, 0x65, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x74,
	0x68, 0x65, 0x72, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x43, 0x68, 0x61, 0x74, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x47, 0x69, 0x66, 0x74, 0x47, 0x69,
	0x76, 0x65, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x10, 0x02, 0x32, 0x75, 0x0a, 0x06, 0x66, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x12, 0x6b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1d, 0x2e,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x61, 0x64, 0x61, 0x72, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1d, 0x22, 0x18, 0x2f, 0x76, 0x31, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x65, 0x6e, 0x75, 0x6d, 0x3a, 0x01,
	0x2a, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64,
	0x65, 0x6d, 0x6f, 0x2d, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x2f, 0x3b, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_friend_friend_proto_rawDescOnce sync.Once
	file_friend_friend_proto_rawDescData = file_friend_friend_proto_rawDesc
)

func file_friend_friend_proto_rawDescGZIP() []byte {
	file_friend_friend_proto_rawDescOnce.Do(func() {
		file_friend_friend_proto_rawDescData = protoimpl.X.CompressGZIP(file_friend_friend_proto_rawDescData)
	})
	return file_friend_friend_proto_rawDescData
}

var file_friend_friend_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_friend_friend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_friend_friend_proto_goTypes = []interface{}{
	(AddFriendType)(0),            // 0: friend.AddFriendType
	(DegreeMethod)(0),             // 1: friend.DegreeMethod
	(*FriendBaseInfo)(nil),        // 2: friend.FriendBaseInfo
	(*RadarSearchPlayerInfo)(nil), // 3: friend.RadarSearchPlayerInfo
}
var file_friend_friend_proto_depIdxs = []int32{
	0, // 0: friend.FriendBaseInfo.addType:type_name -> friend.AddFriendType
	2, // 1: friend.friend.GetFriendInfo:input_type -> friend.FriendBaseInfo
	3, // 2: friend.friend.GetFriendInfo:output_type -> friend.RadarSearchPlayerInfo
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_friend_friend_proto_init() }
func file_friend_friend_proto_init() {
	if File_friend_friend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_friend_friend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendBaseInfo); i {
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
		file_friend_friend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RadarSearchPlayerInfo); i {
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
			RawDescriptor: file_friend_friend_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_friend_friend_proto_goTypes,
		DependencyIndexes: file_friend_friend_proto_depIdxs,
		EnumInfos:         file_friend_friend_proto_enumTypes,
		MessageInfos:      file_friend_friend_proto_msgTypes,
	}.Build()
	File_friend_friend_proto = out.File
	file_friend_friend_proto_rawDesc = nil
	file_friend_friend_proto_goTypes = nil
	file_friend_friend_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FriendClient is the client API for Friend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FriendClient interface {
	GetFriendInfo(ctx context.Context, in *FriendBaseInfo, opts ...grpc.CallOption) (*RadarSearchPlayerInfo, error)
}

type friendClient struct {
	cc grpc.ClientConnInterface
}

func NewFriendClient(cc grpc.ClientConnInterface) FriendClient {
	return &friendClient{cc}
}

func (c *friendClient) GetFriendInfo(ctx context.Context, in *FriendBaseInfo, opts ...grpc.CallOption) (*RadarSearchPlayerInfo, error) {
	out := new(RadarSearchPlayerInfo)
	err := c.cc.Invoke(ctx, "/friend.friend/GetFriendInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FriendServer is the server API for Friend service.
type FriendServer interface {
	GetFriendInfo(context.Context, *FriendBaseInfo) (*RadarSearchPlayerInfo, error)
}

// UnimplementedFriendServer can be embedded to have forward compatible implementations.
type UnimplementedFriendServer struct {
}

func (*UnimplementedFriendServer) GetFriendInfo(context.Context, *FriendBaseInfo) (*RadarSearchPlayerInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriendInfo not implemented")
}

func RegisterFriendServer(s *grpc.Server, srv FriendServer) {
	s.RegisterService(&_Friend_serviceDesc, srv)
}

func _Friend_GetFriendInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FriendBaseInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FriendServer).GetFriendInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/friend.friend/GetFriendInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FriendServer).GetFriendInfo(ctx, req.(*FriendBaseInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Friend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "friend.friend",
	HandlerType: (*FriendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFriendInfo",
			Handler:    _Friend_GetFriendInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "friend/friend.proto",
}
