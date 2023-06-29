// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: login/login.proto

package login

import (
	context "context"
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

type ErrCode int32

const (
	ErrCode_none       ErrCode = 0
	ErrCode_Success    ErrCode = 1
	ErrCode_Unknown    ErrCode = 2
	ErrCode_SessionKey ErrCode = 3
	ErrCode_ImageSize  ErrCode = 4
	ErrCode_FileType   ErrCode = 5
	ErrCode_OpenFile   ErrCode = 6
	ErrCode_ReadFile   ErrCode = 7
	ErrCode_TimeOut    ErrCode = 8
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0: "none",
		1: "Success",
		2: "Unknown",
		3: "SessionKey",
		4: "ImageSize",
		5: "FileType",
		6: "OpenFile",
		7: "ReadFile",
		8: "TimeOut",
	}
	ErrCode_value = map[string]int32{
		"none":       0,
		"Success":    1,
		"Unknown":    2,
		"SessionKey": 3,
		"ImageSize":  4,
		"FileType":   5,
		"OpenFile":   6,
		"ReadFile":   7,
		"TimeOut":    8,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_login_login_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_login_login_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{0}
}

type KickOutReason int32

const (
	KickOutReason_NoReason    KickOutReason = 0
	KickOutReason_RemoteLogin KickOutReason = 1
)

// Enum value maps for KickOutReason.
var (
	KickOutReason_name = map[int32]string{
		0: "NoReason",
		1: "RemoteLogin",
	}
	KickOutReason_value = map[string]int32{
		"NoReason":    0,
		"RemoteLogin": 1,
	}
)

func (x KickOutReason) Enum() *KickOutReason {
	p := new(KickOutReason)
	*p = x
	return p
}

func (x KickOutReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KickOutReason) Descriptor() protoreflect.EnumDescriptor {
	return file_login_login_proto_enumTypes[1].Descriptor()
}

func (KickOutReason) Type() protoreflect.EnumType {
	return &file_login_login_proto_enumTypes[1]
}

func (x KickOutReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KickOutReason.Descriptor instead.
func (KickOutReason) EnumDescriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{1}
}

type LoginData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result            uint32               `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	UserID            uint64               `protobuf:"varint,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Token             string               `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
	SessionID         string               `protobuf:"bytes,4,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	SceneID           int32                `protobuf:"varint,5,opt,name=SceneID,proto3" json:"SceneID,omitempty"`
	CdKey             string               `protobuf:"bytes,6,opt,name=CdKey,proto3" json:"CdKey,omitempty"`
	IP                string               `protobuf:"bytes,7,opt,name=IP,proto3" json:"IP,omitempty"`
	Port              int32                `protobuf:"varint,8,opt,name=Port,proto3" json:"Port,omitempty"`
	ZoneId            int32                `protobuf:"varint,9,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	Reason            string               `protobuf:"bytes,10,opt,name=Reason,proto3" json:"Reason,omitempty"`
	ExpTime           int64                `protobuf:"varint,11,opt,name=ExpTime,proto3" json:"ExpTime,omitempty"`
	BusyLevel         int32                `protobuf:"varint,12,opt,name=BusyLevel,proto3" json:"BusyLevel,omitempty"`
	BusyWaitTime      int32                `protobuf:"varint,13,opt,name=BusyWaitTime,proto3" json:"BusyWaitTime,omitempty"`
	ServerTime        int64                `protobuf:"varint,14,opt,name=ServerTime,proto3" json:"ServerTime,omitempty"`
	RegRegTimeEnd     int64                `protobuf:"varint,15,opt,name=RegRegTimeEnd,proto3" json:"RegRegTimeEnd,omitempty"`
	IsPreReged        bool                 `protobuf:"varint,16,opt,name=IsPreReged,proto3" json:"IsPreReged,omitempty"`
	IsGetShareRewards bool                 `protobuf:"varint,17,opt,name=IsGetShareRewards,proto3" json:"IsGetShareRewards,omitempty"`
	ServerOpenTime    int64                `protobuf:"varint,18,opt,name=ServerOpenTime,proto3" json:"ServerOpenTime,omitempty"`
	IsIpInWhiteList   bool                 `protobuf:"varint,19,opt,name=IsIpInWhiteList,proto3" json:"IsIpInWhiteList,omitempty"`
	ShuShuGameID      string               `protobuf:"bytes,20,opt,name=ShuShuGameID,proto3" json:"ShuShuGameID,omitempty"`
	RegRegTimeStart   int64                `protobuf:"varint,21,opt,name=RegRegTimeStart,proto3" json:"RegRegTimeStart,omitempty"`
	WorldList         []*WorldEndPointInfo `protobuf:"bytes,22,rep,name=WorldList,proto3" json:"WorldList,omitempty"`
	RecommendWorld    []*WorldEndPointInfo `protobuf:"bytes,23,rep,name=RecommendWorld,proto3" json:"RecommendWorld,omitempty"`
	ZoneList          []*ZoneInfo          `protobuf:"bytes,24,rep,name=ZoneList,proto3" json:"ZoneList,omitempty"`
}

func (x *LoginData) Reset() {
	*x = LoginData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginData) ProtoMessage() {}

func (x *LoginData) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginData.ProtoReflect.Descriptor instead.
func (*LoginData) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{0}
}

func (x *LoginData) GetResult() uint32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *LoginData) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LoginData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginData) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *LoginData) GetSceneID() int32 {
	if x != nil {
		return x.SceneID
	}
	return 0
}

func (x *LoginData) GetCdKey() string {
	if x != nil {
		return x.CdKey
	}
	return ""
}

func (x *LoginData) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *LoginData) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *LoginData) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *LoginData) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *LoginData) GetExpTime() int64 {
	if x != nil {
		return x.ExpTime
	}
	return 0
}

func (x *LoginData) GetBusyLevel() int32 {
	if x != nil {
		return x.BusyLevel
	}
	return 0
}

func (x *LoginData) GetBusyWaitTime() int32 {
	if x != nil {
		return x.BusyWaitTime
	}
	return 0
}

func (x *LoginData) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *LoginData) GetRegRegTimeEnd() int64 {
	if x != nil {
		return x.RegRegTimeEnd
	}
	return 0
}

func (x *LoginData) GetIsPreReged() bool {
	if x != nil {
		return x.IsPreReged
	}
	return false
}

func (x *LoginData) GetIsGetShareRewards() bool {
	if x != nil {
		return x.IsGetShareRewards
	}
	return false
}

func (x *LoginData) GetServerOpenTime() int64 {
	if x != nil {
		return x.ServerOpenTime
	}
	return 0
}

func (x *LoginData) GetIsIpInWhiteList() bool {
	if x != nil {
		return x.IsIpInWhiteList
	}
	return false
}

func (x *LoginData) GetShuShuGameID() string {
	if x != nil {
		return x.ShuShuGameID
	}
	return ""
}

func (x *LoginData) GetRegRegTimeStart() int64 {
	if x != nil {
		return x.RegRegTimeStart
	}
	return 0
}

func (x *LoginData) GetWorldList() []*WorldEndPointInfo {
	if x != nil {
		return x.WorldList
	}
	return nil
}

func (x *LoginData) GetRecommendWorld() []*WorldEndPointInfo {
	if x != nil {
		return x.RecommendWorld
	}
	return nil
}

func (x *LoginData) GetZoneList() []*ZoneInfo {
	if x != nil {
		return x.ZoneList
	}
	return nil
}

type ZoneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Status int32 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *ZoneInfo) Reset() {
	*x = ZoneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ZoneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ZoneInfo) ProtoMessage() {}

func (x *ZoneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ZoneInfo.ProtoReflect.Descriptor instead.
func (*ZoneInfo) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{1}
}

func (x *ZoneInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ZoneInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type WorldEndPointInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneId  int32  `protobuf:"varint,1,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	SId     string `protobuf:"bytes,2,opt,name=SId,proto3" json:"SId,omitempty"`
	Addr    string `protobuf:"bytes,3,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	Players int32  `protobuf:"varint,5,opt,name=Players,proto3" json:"Players,omitempty"`
	PIdx    uint32 `protobuf:"varint,6,opt,name=PIdx,proto3" json:"PIdx,omitempty"`
	Max     uint32 `protobuf:"varint,7,opt,name=Max,proto3" json:"Max,omitempty"`
	Stat    int32  `protobuf:"varint,8,opt,name=Stat,proto3" json:"Stat,omitempty"` //0 - 未开服， 1：空闲，2：火热：3：爆满，4：维护中
}

func (x *WorldEndPointInfo) Reset() {
	*x = WorldEndPointInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldEndPointInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldEndPointInfo) ProtoMessage() {}

func (x *WorldEndPointInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldEndPointInfo.ProtoReflect.Descriptor instead.
func (*WorldEndPointInfo) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{2}
}

func (x *WorldEndPointInfo) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *WorldEndPointInfo) GetSId() string {
	if x != nil {
		return x.SId
	}
	return ""
}

func (x *WorldEndPointInfo) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *WorldEndPointInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorldEndPointInfo) GetPlayers() int32 {
	if x != nil {
		return x.Players
	}
	return 0
}

func (x *WorldEndPointInfo) GetPIdx() uint32 {
	if x != nil {
		return x.PIdx
	}
	return 0
}

func (x *WorldEndPointInfo) GetMax() uint32 {
	if x != nil {
		return x.Max
	}
	return 0
}

func (x *WorldEndPointInfo) GetStat() int32 {
	if x != nil {
		return x.Stat
	}
	return 0
}

type WorldInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode       ErrCode      `protobuf:"varint,1,opt,name=ErrCode,proto3,enum=login.ErrCode" json:"ErrCode,omitempty"`
	ZoneId        int32        `protobuf:"varint,2,opt,name=ZoneId,proto3" json:"ZoneId,omitempty"`
	EndPointsInfo []*WorldInfo `protobuf:"bytes,3,rep,name=EndPointsInfo,proto3" json:"EndPointsInfo,omitempty"`
	Recently      []*WorldInfo `protobuf:"bytes,4,rep,name=recently,proto3" json:"recently,omitempty"`
	Recommend     []*WorldInfo `protobuf:"bytes,5,rep,name=recommend,proto3" json:"recommend,omitempty"`
}

func (x *WorldInfo) Reset() {
	*x = WorldInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldInfo) ProtoMessage() {}

func (x *WorldInfo) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldInfo.ProtoReflect.Descriptor instead.
func (*WorldInfo) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{3}
}

func (x *WorldInfo) GetErrCode() ErrCode {
	if x != nil {
		return x.ErrCode
	}
	return ErrCode_none
}

func (x *WorldInfo) GetZoneId() int32 {
	if x != nil {
		return x.ZoneId
	}
	return 0
}

func (x *WorldInfo) GetEndPointsInfo() []*WorldInfo {
	if x != nil {
		return x.EndPointsInfo
	}
	return nil
}

func (x *WorldInfo) GetRecently() []*WorldInfo {
	if x != nil {
		return x.Recently
	}
	return nil
}

func (x *WorldInfo) GetRecommend() []*WorldInfo {
	if x != nil {
		return x.Recommend
	}
	return nil
}

type WorldEndPointInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrCode ErrCode            `protobuf:"varint,1,opt,name=ErrCode,proto3,enum=login.ErrCode" json:"ErrCode,omitempty"`
	Info    *WorldEndPointInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *WorldEndPointInfoRsp) Reset() {
	*x = WorldEndPointInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldEndPointInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldEndPointInfoRsp) ProtoMessage() {}

func (x *WorldEndPointInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldEndPointInfoRsp.ProtoReflect.Descriptor instead.
func (*WorldEndPointInfoRsp) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{4}
}

func (x *WorldEndPointInfoRsp) GetErrCode() ErrCode {
	if x != nil {
		return x.ErrCode
	}
	return ErrCode_none
}

func (x *WorldEndPointInfoRsp) GetInfo() *WorldEndPointInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type KickOutPlayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        uint64        `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	KickOutReason KickOutReason `protobuf:"varint,2,opt,name=KickOutReason,proto3,enum=login.KickOutReason" json:"KickOutReason,omitempty"`
	Reason        string        `protobuf:"bytes,3,opt,name=Reason,proto3" json:"Reason,omitempty"`
	SvrId         string        `protobuf:"bytes,4,opt,name=SvrId,proto3" json:"SvrId,omitempty"`
}

func (x *KickOutPlayer) Reset() {
	*x = KickOutPlayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_login_login_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickOutPlayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickOutPlayer) ProtoMessage() {}

func (x *KickOutPlayer) ProtoReflect() protoreflect.Message {
	mi := &file_login_login_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickOutPlayer.ProtoReflect.Descriptor instead.
func (*KickOutPlayer) Descriptor() ([]byte, []int) {
	return file_login_login_proto_rawDescGZIP(), []int{5}
}

func (x *KickOutPlayer) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *KickOutPlayer) GetKickOutReason() KickOutReason {
	if x != nil {
		return x.KickOutReason
	}
	return KickOutReason_NoReason
}

func (x *KickOutPlayer) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *KickOutPlayer) GetSvrId() string {
	if x != nil {
		return x.SvrId
	}
	return ""
}

var File_login_login_proto protoreflect.FileDescriptor

var file_login_login_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0xaa, 0x06, 0x0a, 0x09, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x64, 0x4b, 0x65, 0x79, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x50, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04,
	0x50, 0x6f, 0x72, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x45, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x45, 0x78, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x75,
	0x73, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x42,
	0x75, 0x73, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x42, 0x75, 0x73, 0x79,
	0x57, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x42, 0x75, 0x73, 0x79, 0x57, 0x61, 0x69, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x52, 0x65, 0x67, 0x52, 0x65, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x52, 0x65, 0x67, 0x52, 0x65, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x45,
	0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x73, 0x50, 0x72, 0x65, 0x52, 0x65, 0x67, 0x65, 0x64,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x49, 0x73, 0x50, 0x72, 0x65, 0x52, 0x65, 0x67,
	0x65, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x49, 0x73, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x49,
	0x73, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73,
	0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x49, 0x73, 0x49, 0x70,
	0x49, 0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0f, 0x49, 0x73, 0x49, 0x70, 0x49, 0x6e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x68, 0x75, 0x53, 0x68, 0x75, 0x47, 0x61, 0x6d, 0x65,
	0x49, 0x44, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x68, 0x75, 0x53, 0x68, 0x75,
	0x47, 0x61, 0x6d, 0x65, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x52, 0x65, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x52, 0x65, 0x67, 0x52, 0x65, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x36, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x16, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x57,
	0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x18, 0x17, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x6e,
	0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x52, 0x65, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x5a, 0x6f,
	0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x5a,
	0x6f, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x32, 0x0a, 0x08, 0x5a, 0x6f, 0x6e, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xb9, 0x01, 0x0a, 0x11,
	0x57, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x49, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x41,
	0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x50, 0x49, 0x64, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x49, 0x64,
	0x78, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x61, 0x78, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03,
	0x4d, 0x61, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x53, 0x74, 0x61, 0x74, 0x22, 0xe3, 0x01, 0x0a, 0x09, 0x57, 0x6f, 0x72, 0x6c,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x0d, 0x45, 0x6e, 0x64, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x0d, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2c, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x6c, 0x79, 0x12, 0x2e, 0x0a,
	0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x22, 0x6e, 0x0a,
	0x14, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x28, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x6e, 0x64, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x91, 0x01,
	0x0a, 0x0d, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x3a, 0x0a, 0x0d, 0x4b, 0x69, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x0d, 0x4b, 0x69, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x53,
	0x76, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x76, 0x72, 0x49,
	0x64, 0x2a, 0x83, 0x01, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x08, 0x0a,
	0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x10,
	0x03, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x10, 0x04,
	0x12, 0x0c, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x05, 0x12, 0x0c,
	0x0a, 0x08, 0x4f, 0x70, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x06, 0x12, 0x0c, 0x0a, 0x08,
	0x52, 0x65, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x69,
	0x6d, 0x65, 0x4f, 0x75, 0x74, 0x10, 0x08, 0x2a, 0x2e, 0x0a, 0x0d, 0x4b, 0x69, 0x63, 0x6b, 0x4f,
	0x75, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x32, 0x3f, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x12, 0x36, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x1b, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x45, 0x6e, 0x64, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x72, 0x70, 0x63,
	0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x32, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x3b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_login_login_proto_rawDescOnce sync.Once
	file_login_login_proto_rawDescData = file_login_login_proto_rawDesc
)

func file_login_login_proto_rawDescGZIP() []byte {
	file_login_login_proto_rawDescOnce.Do(func() {
		file_login_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_login_login_proto_rawDescData)
	})
	return file_login_login_proto_rawDescData
}

var file_login_login_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_login_login_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_login_login_proto_goTypes = []interface{}{
	(ErrCode)(0),                 // 0: login.ErrCode
	(KickOutReason)(0),           // 1: login.KickOutReason
	(*LoginData)(nil),            // 2: login.LoginData
	(*ZoneInfo)(nil),             // 3: login.ZoneInfo
	(*WorldEndPointInfo)(nil),    // 4: login.WorldEndPointInfo
	(*WorldInfo)(nil),            // 5: login.WorldInfo
	(*WorldEndPointInfoRsp)(nil), // 6: login.WorldEndPointInfoRsp
	(*KickOutPlayer)(nil),        // 7: login.KickOutPlayer
}
var file_login_login_proto_depIdxs = []int32{
	4,  // 0: login.LoginData.WorldList:type_name -> login.WorldEndPointInfo
	4,  // 1: login.LoginData.RecommendWorld:type_name -> login.WorldEndPointInfo
	3,  // 2: login.LoginData.ZoneList:type_name -> login.ZoneInfo
	0,  // 3: login.WorldInfo.ErrCode:type_name -> login.ErrCode
	5,  // 4: login.WorldInfo.EndPointsInfo:type_name -> login.WorldInfo
	5,  // 5: login.WorldInfo.recently:type_name -> login.WorldInfo
	5,  // 6: login.WorldInfo.recommend:type_name -> login.WorldInfo
	0,  // 7: login.WorldEndPointInfoRsp.ErrCode:type_name -> login.ErrCode
	4,  // 8: login.WorldEndPointInfoRsp.info:type_name -> login.WorldEndPointInfo
	1,  // 9: login.KickOutPlayer.KickOutReason:type_name -> login.KickOutReason
	2,  // 10: login.login.login:input_type -> login.LoginData
	6,  // 11: login.login.login:output_type -> login.WorldEndPointInfoRsp
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_login_login_proto_init() }
func file_login_login_proto_init() {
	if File_login_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_login_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginData); i {
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
		file_login_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ZoneInfo); i {
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
		file_login_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldEndPointInfo); i {
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
		file_login_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldInfo); i {
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
		file_login_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldEndPointInfoRsp); i {
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
		file_login_login_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickOutPlayer); i {
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
			RawDescriptor: file_login_login_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_login_login_proto_goTypes,
		DependencyIndexes: file_login_login_proto_depIdxs,
		EnumInfos:         file_login_login_proto_enumTypes,
		MessageInfos:      file_login_login_proto_msgTypes,
	}.Build()
	File_login_login_proto = out.File
	file_login_login_proto_rawDesc = nil
	file_login_login_proto_goTypes = nil
	file_login_login_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LoginClient is the client API for Login service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginClient interface {
	Login(ctx context.Context, in *LoginData, opts ...grpc.CallOption) (*WorldEndPointInfoRsp, error)
}

type loginClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginClient(cc grpc.ClientConnInterface) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginData, opts ...grpc.CallOption) (*WorldEndPointInfoRsp, error) {
	out := new(WorldEndPointInfoRsp)
	err := c.cc.Invoke(ctx, "/login.login/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServer is the server API for Login service.
type LoginServer interface {
	Login(context.Context, *LoginData) (*WorldEndPointInfoRsp, error)
}

// UnimplementedLoginServer can be embedded to have forward compatible implementations.
type UnimplementedLoginServer struct {
}

func (*UnimplementedLoginServer) Login(context.Context, *LoginData) (*WorldEndPointInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/login.login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "login.login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login/login.proto",
}
