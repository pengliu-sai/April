// Code generated by protoc-gen-go.
// source: src/protos/game/game.proto
// DO NOT EDIT!

/*
Package game is a generated protocol buffer package.

It is generated from these files:
	src/protos/game/game.proto

It has these top-level messages:
	Game_PingC2S
	Game_PingS2C
	Game_RegisterUserIDC2S
	Game_RegisterUserIDS2C
	Game_RegisterRoleIDC2S
	Game_RegisterRoleIDS2C
	Game_RandomRoleNameC2S
	Game_RandomRoleNameS2C
	Game_RoleCreateC2S
	Game_RoleCreateS2C
	RoleInfo
	Game_RoleInfoListC2S
	Game_RoleInfoListS2C
	Game_EnterScenesC2S
	Game_EnterScenesS2C
	Game_ExitScenesC2S
	Game_ExitScenesS2C
	Game_SendChatC2S
	Game_SendChatS2C
	Game_Receive_ChatS2C
	Game_RoleInfoByRoleIDC2S
	Game_RoleInfoByRoleIDS2C
	Game_AddFriendC2S
	Game_AddFriendS2C
	Game_FriendListC2S
	Game_FriendListS2C
*/
package game

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SexEnum int32

const (
	SexEnum_male   SexEnum = 0
	SexEnum_female SexEnum = 1
)

var SexEnum_name = map[int32]string{
	0: "male",
	1: "female",
}
var SexEnum_value = map[string]int32{
	"male":   0,
	"female": 1,
}

func (x SexEnum) Enum() *SexEnum {
	p := new(SexEnum)
	*p = x
	return p
}
func (x SexEnum) String() string {
	return proto.EnumName(SexEnum_name, int32(x))
}
func (x *SexEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SexEnum_value, data, "SexEnum")
	if err != nil {
		return err
	}
	*x = SexEnum(value)
	return nil
}
func (SexEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RacesEnum int32

const (
	RacesEnum_warrior   RacesEnum = 0
	RacesEnum_enchanter RacesEnum = 1
	RacesEnum_archer    RacesEnum = 2
)

var RacesEnum_name = map[int32]string{
	0: "warrior",
	1: "enchanter",
	2: "archer",
}
var RacesEnum_value = map[string]int32{
	"warrior":   0,
	"enchanter": 1,
	"archer":    2,
}

func (x RacesEnum) Enum() *RacesEnum {
	p := new(RacesEnum)
	*p = x
	return p
}
func (x RacesEnum) String() string {
	return proto.EnumName(RacesEnum_name, int32(x))
}
func (x *RacesEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RacesEnum_value, data, "RacesEnum")
	if err != nil {
		return err
	}
	*x = RacesEnum(value)
	return nil
}
func (RacesEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Game_PingC2S struct {
	Content          *string `protobuf:"bytes,1,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_PingC2S) Reset()                    { *m = Game_PingC2S{} }
func (m *Game_PingC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_PingC2S) ProtoMessage()               {}
func (*Game_PingC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Game_PingC2S) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type Game_PingS2C struct {
	Content          *string `protobuf:"bytes,1,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_PingS2C) Reset()                    { *m = Game_PingS2C{} }
func (m *Game_PingS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_PingS2C) ProtoMessage()               {}
func (*Game_PingS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Game_PingS2C) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type Game_RegisterUserIDC2S struct {
	UserID           *uint64 `protobuf:"varint,1,req,name=UserID,json=userID" json:"UserID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_RegisterUserIDC2S) Reset()                    { *m = Game_RegisterUserIDC2S{} }
func (m *Game_RegisterUserIDC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_RegisterUserIDC2S) ProtoMessage()               {}
func (*Game_RegisterUserIDC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Game_RegisterUserIDC2S) GetUserID() uint64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

type Game_RegisterUserIDS2C struct {
	Result           *bool   `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	UserID           *uint64 `protobuf:"varint,2,req,name=UserID,json=userID" json:"UserID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_RegisterUserIDS2C) Reset()                    { *m = Game_RegisterUserIDS2C{} }
func (m *Game_RegisterUserIDS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_RegisterUserIDS2C) ProtoMessage()               {}
func (*Game_RegisterUserIDS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Game_RegisterUserIDS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

func (m *Game_RegisterUserIDS2C) GetUserID() uint64 {
	if m != nil && m.UserID != nil {
		return *m.UserID
	}
	return 0
}

type Game_RegisterRoleIDC2S struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_RegisterRoleIDC2S) Reset()                    { *m = Game_RegisterRoleIDC2S{} }
func (m *Game_RegisterRoleIDC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_RegisterRoleIDC2S) ProtoMessage()               {}
func (*Game_RegisterRoleIDC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Game_RegisterRoleIDC2S) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type Game_RegisterRoleIDS2C struct {
	Result           *bool  `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_RegisterRoleIDS2C) Reset()                    { *m = Game_RegisterRoleIDS2C{} }
func (m *Game_RegisterRoleIDS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_RegisterRoleIDS2C) ProtoMessage()               {}
func (*Game_RegisterRoleIDS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Game_RegisterRoleIDS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type Game_RandomRoleNameC2S struct {
	Sex              *SexEnum `protobuf:"varint,1,req,name=Sex,json=sex,enum=SexEnum" json:"Sex,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Game_RandomRoleNameC2S) Reset()                    { *m = Game_RandomRoleNameC2S{} }
func (m *Game_RandomRoleNameC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_RandomRoleNameC2S) ProtoMessage()               {}
func (*Game_RandomRoleNameC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Game_RandomRoleNameC2S) GetSex() SexEnum {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return SexEnum_male
}

type Game_RandomRoleNameS2C struct {
	Name             *string `protobuf:"bytes,1,req,name=Name,json=name" json:"Name,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_RandomRoleNameS2C) Reset()                    { *m = Game_RandomRoleNameS2C{} }
func (m *Game_RandomRoleNameS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_RandomRoleNameS2C) ProtoMessage()               {}
func (*Game_RandomRoleNameS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Game_RandomRoleNameS2C) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

type Game_RoleCreateC2S struct {
	Name             *string    `protobuf:"bytes,1,req,name=Name,json=name" json:"Name,omitempty"`
	Sex              *SexEnum   `protobuf:"varint,2,req,name=Sex,json=sex,enum=SexEnum" json:"Sex,omitempty"`
	Race             *RacesEnum `protobuf:"varint,3,req,name=race,enum=RacesEnum" json:"race,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Game_RoleCreateC2S) Reset()                    { *m = Game_RoleCreateC2S{} }
func (m *Game_RoleCreateC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_RoleCreateC2S) ProtoMessage()               {}
func (*Game_RoleCreateC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Game_RoleCreateC2S) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Game_RoleCreateC2S) GetSex() SexEnum {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return SexEnum_male
}

func (m *Game_RoleCreateC2S) GetRace() RacesEnum {
	if m != nil && m.Race != nil {
		return *m.Race
	}
	return RacesEnum_warrior
}

type Game_RoleCreateS2C struct {
	RoleID           *uint64    `protobuf:"varint,1,opt,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	RoleName         *string    `protobuf:"bytes,2,opt,name=RoleName,json=roleName" json:"RoleName,omitempty"`
	Sex              *SexEnum   `protobuf:"varint,3,opt,name=Sex,json=sex,enum=SexEnum" json:"Sex,omitempty"`
	Race             *RacesEnum `protobuf:"varint,4,opt,name=race,enum=RacesEnum" json:"race,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *Game_RoleCreateS2C) Reset()                    { *m = Game_RoleCreateS2C{} }
func (m *Game_RoleCreateS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_RoleCreateS2C) ProtoMessage()               {}
func (*Game_RoleCreateS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Game_RoleCreateS2C) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *Game_RoleCreateS2C) GetRoleName() string {
	if m != nil && m.RoleName != nil {
		return *m.RoleName
	}
	return ""
}

func (m *Game_RoleCreateS2C) GetSex() SexEnum {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return SexEnum_male
}

func (m *Game_RoleCreateS2C) GetRace() RacesEnum {
	if m != nil && m.Race != nil {
		return *m.Race
	}
	return RacesEnum_warrior
}

type RoleInfo struct {
	RoleID           *uint64    `protobuf:"varint,1,opt,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	RoleName         *string    `protobuf:"bytes,2,opt,name=RoleName,json=roleName" json:"RoleName,omitempty"`
	Sex              *SexEnum   `protobuf:"varint,3,opt,name=sex,enum=SexEnum" json:"sex,omitempty"`
	Race             *RacesEnum `protobuf:"varint,4,opt,name=race,enum=RacesEnum" json:"race,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RoleInfo) Reset()                    { *m = RoleInfo{} }
func (m *RoleInfo) String() string            { return proto.CompactTextString(m) }
func (*RoleInfo) ProtoMessage()               {}
func (*RoleInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RoleInfo) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *RoleInfo) GetRoleName() string {
	if m != nil && m.RoleName != nil {
		return *m.RoleName
	}
	return ""
}

func (m *RoleInfo) GetSex() SexEnum {
	if m != nil && m.Sex != nil {
		return *m.Sex
	}
	return SexEnum_male
}

func (m *RoleInfo) GetRace() RacesEnum {
	if m != nil && m.Race != nil {
		return *m.Race
	}
	return RacesEnum_warrior
}

type Game_RoleInfoListC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_RoleInfoListC2S) Reset()                    { *m = Game_RoleInfoListC2S{} }
func (m *Game_RoleInfoListC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_RoleInfoListC2S) ProtoMessage()               {}
func (*Game_RoleInfoListC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type Game_RoleInfoListS2C struct {
	Roles            []*RoleInfo `protobuf:"bytes,1,rep,name=roles" json:"roles,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Game_RoleInfoListS2C) Reset()                    { *m = Game_RoleInfoListS2C{} }
func (m *Game_RoleInfoListS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_RoleInfoListS2C) ProtoMessage()               {}
func (*Game_RoleInfoListS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Game_RoleInfoListS2C) GetRoles() []*RoleInfo {
	if m != nil {
		return m.Roles
	}
	return nil
}

type Game_EnterScenesC2S struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_EnterScenesC2S) Reset()                    { *m = Game_EnterScenesC2S{} }
func (m *Game_EnterScenesC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_EnterScenesC2S) ProtoMessage()               {}
func (*Game_EnterScenesC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Game_EnterScenesC2S) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type Game_EnterScenesS2C struct {
	Result           *bool  `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_EnterScenesS2C) Reset()                    { *m = Game_EnterScenesS2C{} }
func (m *Game_EnterScenesS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_EnterScenesS2C) ProtoMessage()               {}
func (*Game_EnterScenesS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *Game_EnterScenesS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type Game_ExitScenesC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_ExitScenesC2S) Reset()                    { *m = Game_ExitScenesC2S{} }
func (m *Game_ExitScenesC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_ExitScenesC2S) ProtoMessage()               {}
func (*Game_ExitScenesC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type Game_ExitScenesS2C struct {
	Result           *bool  `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_ExitScenesS2C) Reset()                    { *m = Game_ExitScenesS2C{} }
func (m *Game_ExitScenesS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_ExitScenesS2C) ProtoMessage()               {}
func (*Game_ExitScenesS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Game_ExitScenesS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type Game_SendChatC2S struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	Content          *string `protobuf:"bytes,2,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_SendChatC2S) Reset()                    { *m = Game_SendChatC2S{} }
func (m *Game_SendChatC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_SendChatC2S) ProtoMessage()               {}
func (*Game_SendChatC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *Game_SendChatC2S) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *Game_SendChatC2S) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type Game_SendChatS2C struct {
	Result           *bool  `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_SendChatS2C) Reset()                    { *m = Game_SendChatS2C{} }
func (m *Game_SendChatS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_SendChatS2C) ProtoMessage()               {}
func (*Game_SendChatS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *Game_SendChatS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

// 接收聊天信息
type Game_Receive_ChatS2C struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	RoleName         *string `protobuf:"bytes,2,req,name=RoleName,json=roleName" json:"RoleName,omitempty"`
	Content          *string `protobuf:"bytes,3,req,name=Content,json=content" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_Receive_ChatS2C) Reset()                    { *m = Game_Receive_ChatS2C{} }
func (m *Game_Receive_ChatS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_Receive_ChatS2C) ProtoMessage()               {}
func (*Game_Receive_ChatS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *Game_Receive_ChatS2C) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

func (m *Game_Receive_ChatS2C) GetRoleName() string {
	if m != nil && m.RoleName != nil {
		return *m.RoleName
	}
	return ""
}

func (m *Game_Receive_ChatS2C) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

type Game_RoleInfoByRoleIDC2S struct {
	RoleID           *uint64 `protobuf:"varint,1,req,name=RoleID,json=roleID" json:"RoleID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_RoleInfoByRoleIDC2S) Reset()                    { *m = Game_RoleInfoByRoleIDC2S{} }
func (m *Game_RoleInfoByRoleIDC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_RoleInfoByRoleIDC2S) ProtoMessage()               {}
func (*Game_RoleInfoByRoleIDC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Game_RoleInfoByRoleIDC2S) GetRoleID() uint64 {
	if m != nil && m.RoleID != nil {
		return *m.RoleID
	}
	return 0
}

type Game_RoleInfoByRoleIDS2C struct {
	Role             *RoleInfo `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Game_RoleInfoByRoleIDS2C) Reset()                    { *m = Game_RoleInfoByRoleIDS2C{} }
func (m *Game_RoleInfoByRoleIDS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_RoleInfoByRoleIDS2C) ProtoMessage()               {}
func (*Game_RoleInfoByRoleIDS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *Game_RoleInfoByRoleIDS2C) GetRole() *RoleInfo {
	if m != nil {
		return m.Role
	}
	return nil
}

type Game_AddFriendC2S struct {
	FriendID         *uint64 `protobuf:"varint,1,req,name=FriendID,json=friendID" json:"FriendID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Game_AddFriendC2S) Reset()                    { *m = Game_AddFriendC2S{} }
func (m *Game_AddFriendC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_AddFriendC2S) ProtoMessage()               {}
func (*Game_AddFriendC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *Game_AddFriendC2S) GetFriendID() uint64 {
	if m != nil && m.FriendID != nil {
		return *m.FriendID
	}
	return 0
}

type Game_AddFriendS2C struct {
	Result           *bool  `protobuf:"varint,1,req,name=Result,json=result" json:"Result,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_AddFriendS2C) Reset()                    { *m = Game_AddFriendS2C{} }
func (m *Game_AddFriendS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_AddFriendS2C) ProtoMessage()               {}
func (*Game_AddFriendS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *Game_AddFriendS2C) GetResult() bool {
	if m != nil && m.Result != nil {
		return *m.Result
	}
	return false
}

type Game_FriendListC2S struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *Game_FriendListC2S) Reset()                    { *m = Game_FriendListC2S{} }
func (m *Game_FriendListC2S) String() string            { return proto.CompactTextString(m) }
func (*Game_FriendListC2S) ProtoMessage()               {}
func (*Game_FriendListC2S) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

type Game_FriendListS2C struct {
	Friends          []*RoleInfo `protobuf:"bytes,1,rep,name=friends" json:"friends,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Game_FriendListS2C) Reset()                    { *m = Game_FriendListS2C{} }
func (m *Game_FriendListS2C) String() string            { return proto.CompactTextString(m) }
func (*Game_FriendListS2C) ProtoMessage()               {}
func (*Game_FriendListS2C) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *Game_FriendListS2C) GetFriends() []*RoleInfo {
	if m != nil {
		return m.Friends
	}
	return nil
}

func init() {
	proto.RegisterType((*Game_PingC2S)(nil), "Game_PingC2S")
	proto.RegisterType((*Game_PingS2C)(nil), "Game_PingS2C")
	proto.RegisterType((*Game_RegisterUserIDC2S)(nil), "Game_RegisterUserIDC2S")
	proto.RegisterType((*Game_RegisterUserIDS2C)(nil), "Game_RegisterUserIDS2C")
	proto.RegisterType((*Game_RegisterRoleIDC2S)(nil), "Game_RegisterRoleIDC2S")
	proto.RegisterType((*Game_RegisterRoleIDS2C)(nil), "Game_RegisterRoleIDS2C")
	proto.RegisterType((*Game_RandomRoleNameC2S)(nil), "Game_RandomRoleNameC2S")
	proto.RegisterType((*Game_RandomRoleNameS2C)(nil), "Game_RandomRoleNameS2C")
	proto.RegisterType((*Game_RoleCreateC2S)(nil), "Game_RoleCreateC2S")
	proto.RegisterType((*Game_RoleCreateS2C)(nil), "Game_RoleCreateS2C")
	proto.RegisterType((*RoleInfo)(nil), "RoleInfo")
	proto.RegisterType((*Game_RoleInfoListC2S)(nil), "Game_RoleInfoListC2S")
	proto.RegisterType((*Game_RoleInfoListS2C)(nil), "Game_RoleInfoListS2C")
	proto.RegisterType((*Game_EnterScenesC2S)(nil), "Game_EnterScenesC2S")
	proto.RegisterType((*Game_EnterScenesS2C)(nil), "Game_EnterScenesS2C")
	proto.RegisterType((*Game_ExitScenesC2S)(nil), "Game_ExitScenesC2S")
	proto.RegisterType((*Game_ExitScenesS2C)(nil), "Game_ExitScenesS2C")
	proto.RegisterType((*Game_SendChatC2S)(nil), "Game_SendChatC2S")
	proto.RegisterType((*Game_SendChatS2C)(nil), "Game_SendChatS2C")
	proto.RegisterType((*Game_Receive_ChatS2C)(nil), "Game_Receive_ChatS2C")
	proto.RegisterType((*Game_RoleInfoByRoleIDC2S)(nil), "Game_RoleInfoByRoleIDC2S")
	proto.RegisterType((*Game_RoleInfoByRoleIDS2C)(nil), "Game_RoleInfoByRoleIDS2C")
	proto.RegisterType((*Game_AddFriendC2S)(nil), "Game_AddFriendC2S")
	proto.RegisterType((*Game_AddFriendS2C)(nil), "Game_AddFriendS2C")
	proto.RegisterType((*Game_FriendListC2S)(nil), "Game_FriendListC2S")
	proto.RegisterType((*Game_FriendListS2C)(nil), "Game_FriendListS2C")
	proto.RegisterEnum("SexEnum", SexEnum_name, SexEnum_value)
	proto.RegisterEnum("RacesEnum", RacesEnum_name, RacesEnum_value)
}

var fileDescriptor0 = []byte{
	// 557 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x8f, 0xd2, 0x40,
	0x10, 0xb6, 0x3f, 0x84, 0x32, 0xa7, 0x06, 0xeb, 0x85, 0x90, 0x4b, 0xf4, 0x4c, 0x7d, 0x21, 0xa8,
	0x60, 0xd0, 0xc4, 0xdc, 0xa3, 0x72, 0xa8, 0x97, 0x18, 0x63, 0x4a, 0x7c, 0xbe, 0x34, 0xed, 0x00,
	0x4d, 0xa0, 0x35, 0xdb, 0xa2, 0xe8, 0xb3, 0x7f, 0xb8, 0xb3, 0xb3, 0x6c, 0xa5, 0x50, 0xea, 0xe5,
	0x5e, 0xc8, 0x7e, 0xb3, 0xdf, 0xcc, 0xf7, 0x31, 0x3b, 0x53, 0x38, 0xcb, 0x44, 0x38, 0xfc, 0x2e,
	0xd2, 0x3c, 0xcd, 0x86, 0xf3, 0x60, 0x85, 0xfc, 0x33, 0xe0, 0x80, 0xd7, 0x83, 0x7b, 0x1f, 0x09,
	0x5d, 0x7f, 0x8d, 0x93, 0xf9, 0x78, 0x34, 0x75, 0xbb, 0xd0, 0x1c, 0xa7, 0x49, 0x8e, 0x49, 0xde,
	0x35, 0x9e, 0x9a, 0xbd, 0x96, 0xdf, 0x0c, 0x15, 0x2c, 0x31, 0xa7, 0xa3, 0x71, 0x0d, 0xf3, 0x15,
	0x74, 0x98, 0xe9, 0xe3, 0x3c, 0xce, 0x72, 0x14, 0xdf, 0x32, 0x14, 0x57, 0x97, 0xb2, 0x7a, 0x07,
	0x1a, 0x0a, 0x70, 0x8a, 0xed, 0x37, 0xd6, 0x8c, 0xbc, 0x4f, 0x95, 0x19, 0x52, 0x85, 0x32, 0x7c,
	0xcc, 0xd6, 0x4b, 0x25, 0xe2, 0xf8, 0x0d, 0xc1, 0x68, 0xa7, 0x92, 0x59, 0xaa, 0xb4, 0xaf, 0xed,
	0xa7, 0x4b, 0x2c, 0xb4, 0x15, 0xd0, 0xda, 0x82, 0xd1, 0x91, 0x8c, 0x1a, 0x6d, 0xef, 0x8d, 0xce,
	0x08, 0x92, 0x28, 0x5d, 0x49, 0xfe, 0x17, 0x82, 0x52, 0xe3, 0x0c, 0xac, 0x29, 0x6e, 0x98, 0xfe,
	0x60, 0xe4, 0x0c, 0xe8, 0x3c, 0x49, 0xd6, 0x2b, 0xdf, 0xca, 0x70, 0xe3, 0xbd, 0xa8, 0xcc, 0x92,
	0x3a, 0x2e, 0xd8, 0xf2, 0xb8, 0x6d, 0xa3, 0x9d, 0xd0, 0xd9, 0x8b, 0xc0, 0x55, 0x6c, 0xe2, 0x8d,
	0x05, 0x06, 0x39, 0xd7, 0xaf, 0x60, 0x6a, 0x4d, 0xb3, 0x42, 0xd3, 0x7d, 0x02, 0xb6, 0x08, 0x42,
	0xec, 0x5a, 0x7c, 0x09, 0x03, 0x9f, 0x40, 0xc6, 0xd7, 0x1c, 0xf7, 0xfe, 0x18, 0x07, 0x32, 0xfa,
	0x8f, 0xeb, 0x56, 0x19, 0xff, 0x5a, 0x45, 0x52, 0x8e, 0xf6, 0x4d, 0x7a, 0x06, 0x59, 0x70, 0xc4,
	0x16, 0x6b, 0x1b, 0x16, 0x85, 0x8f, 0xda, 0xb0, 0xf9, 0xf2, 0xd0, 0xc6, 0x6f, 0x55, 0xf7, 0x2a,
	0x99, 0xa5, 0xb7, 0xd5, 0xce, 0x6e, 0xab, 0xdd, 0x81, 0xd3, 0xa2, 0x03, 0xd2, 0xc0, 0x67, 0x1a,
	0x02, 0x6a, 0xb5, 0xf7, 0xb6, 0x22, 0x2e, 0x7b, 0x73, 0x0e, 0x77, 0xa5, 0x6e, 0x46, 0xf6, 0xac,
	0xde, 0xc9, 0xa8, 0x35, 0xd0, 0x04, 0x5f, 0xc5, 0xbd, 0x97, 0xf0, 0x88, 0x13, 0x27, 0xb4, 0x0c,
	0x62, 0x1a, 0x62, 0x82, 0x59, 0xdd, 0xf8, 0x55, 0xd0, 0xeb, 0x66, 0xef, 0x74, 0xfb, 0x60, 0x93,
	0x4d, 0x9c, 0x17, 0xc5, 0x69, 0xb6, 0xf6, 0xa3, 0x75, 0x35, 0x2e, 0xa1, 0xcd, 0xec, 0x29, 0x26,
	0xd1, 0x78, 0x11, 0xe4, 0x35, 0xf6, 0x76, 0xb7, 0xdc, 0x2c, 0x6f, 0x79, 0x7f, 0xaf, 0x4a, 0x9d,
	0x62, 0xa4, 0x9b, 0x89, 0x21, 0xc6, 0x3f, 0xf0, 0x7a, 0x97, 0x5f, 0xa5, 0x5a, 0x7e, 0x6c, 0xb3,
	0xf4, 0xd8, 0x3b, 0x8e, 0xac, 0xb2, 0xa3, 0x11, 0x74, 0x4b, 0x4f, 0xf6, 0xfe, 0xd7, 0xff, 0xb7,
	0xff, 0xe2, 0x48, 0x8e, 0x74, 0xf7, 0x98, 0x46, 0x87, 0x00, 0x0f, 0x62, 0xe9, 0xa5, 0x39, 0xec,
	0x0d, 0xe1, 0x21, 0xa7, 0xbe, 0x8b, 0xa2, 0x0f, 0x22, 0x96, 0x5d, 0xe0, 0x2f, 0x80, 0xa3, 0x40,
	0xa1, 0xe4, 0xcc, 0xb6, 0xd8, 0x7b, 0xbe, 0x9f, 0x70, 0x93, 0x87, 0x56, 0x4c, 0x3d, 0x95, 0x17,
	0x07, 0x51, 0x59, 0xe3, 0x19, 0x34, 0x95, 0x48, 0xc5, 0x54, 0xea, 0x9b, 0xfe, 0x39, 0x34, 0xb7,
	0x8b, 0xe1, 0x3a, 0x60, 0xaf, 0x82, 0x25, 0xb6, 0xef, 0xb8, 0x00, 0x8d, 0x19, 0xf2, 0xd9, 0xe8,
	0xbf, 0x86, 0x56, 0xb1, 0x1c, 0xee, 0x09, 0x34, 0x7f, 0x06, 0x42, 0xc4, 0xa9, 0x20, 0xd6, 0x7d,
	0x68, 0x61, 0x12, 0x2e, 0x02, 0x39, 0xa2, 0x6d, 0x43, 0x26, 0x05, 0x22, 0x5c, 0xd0, 0xd9, 0xfc,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x70, 0x5c, 0x88, 0xa6, 0x5c, 0x06, 0x00, 0x00,
}
