// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: work/sg/common/v1/api.proto

// package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

package v1

import (
	fmt "fmt"
	_ "github.com/bilibili/kratos/tool/protobuf/pkg/extensions/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SgMessageType int32

const (
	SgMessageType_DMT_UNKNOWN SgMessageType = 0
	// 待处理
	SgMessageType_DMT_TBD SgMessageType = 1
	// 动态
	SgMessageType_DMT_TRENDS SgMessageType = 2
)

var SgMessageType_name = map[int32]string{
	0: "DMT_UNKNOWN",
	1: "DMT_TBD",
	2: "DMT_TRENDS",
}

var SgMessageType_value = map[string]int32{
	"DMT_UNKNOWN": 0,
	"DMT_TBD":     1,
	"DMT_TRENDS":  2,
}

func (x SgMessageType) String() string {
	return proto.EnumName(SgMessageType_name, int32(x))
}

func (SgMessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94955c5bcc4b00ce, []int{0}
}

// 消息类型
type SgEvent int32

const (
	SgEvent_SE_UNKNOWN SgEvent = 0
	// 添加协作者
	//（某用户被添加为协作者）
	SgEvent_SE_ADD_AS_COLL SgEvent = 1
	// 移除协作者
	//（协作者被移除协作身份）
	SgEvent_SE_DEL_AS_COLL SgEvent = 2
	// 权限修改
	//（协作者的协作权限被操作人修改）
	SgEvent_SE_CHG_PEM SgEvent = 3
	// 转让所有权
	//（所有者将文件的所有权转让给受让人）
	SgEvent_SE_TSF SgEvent = 4
	// 申请访问权限
	//（申请人申请访问文件（夹））
	SgEvent_SE_APP_AS_COLL SgEvent = 5
	// 评论提醒
	//（评论人在文档中发表了评论、@到了当前用户）
	SgEvent_SE_NT_CMENT SgEvent = 6
	// 修订提醒
	//（修订人在文档中进行了修订）
	SgEvent_SE_NT_REVISE SgEvent = 7
	// 文档因为含有违规内容被系统关闭公开分享
	SgEvent_SE_NT_COLSE_PUBLIC_SHARE SgEvent = 8
)

var SgEvent_name = map[int32]string{
	0: "SE_UNKNOWN",
	1: "SE_ADD_AS_COLL",
	2: "SE_DEL_AS_COLL",
	3: "SE_CHG_PEM",
	4: "SE_TSF",
	5: "SE_APP_AS_COLL",
	6: "SE_NT_CMENT",
	7: "SE_NT_REVISE",
	8: "SE_NT_COLSE_PUBLIC_SHARE",
}

var SgEvent_value = map[string]int32{
	"SE_UNKNOWN":               0,
	"SE_ADD_AS_COLL":           1,
	"SE_DEL_AS_COLL":           2,
	"SE_CHG_PEM":               3,
	"SE_TSF":                   4,
	"SE_APP_AS_COLL":           5,
	"SE_NT_CMENT":              6,
	"SE_NT_REVISE":             7,
	"SE_NT_COLSE_PUBLIC_SHARE": 8,
}

func (x SgEvent) String() string {
	return proto.EnumName(SgEvent_name, int32(x))
}

func (SgEvent) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94955c5bcc4b00ce, []int{1}
}

// 消息的权重登记
type SgMessagePriority int32

const (
	SgMessagePriority_SMP_UNKNOWN         SgMessagePriority = 0
	SgMessagePriority_SMP_VERY_LOW_VALUE  SgMessagePriority = 1
	SgMessagePriority_SMP_LOW_VALUE       SgMessagePriority = 2
	SgMessagePriority_SMP_NORMAL_VALUE    SgMessagePriority = 3
	SgMessagePriority_SMP_HIGH_VALUE      SgMessagePriority = 4
	SgMessagePriority_SMP_VERY_HIGH_VALUE SgMessagePriority = 5
)

var SgMessagePriority_name = map[int32]string{
	0: "SMP_UNKNOWN",
	1: "SMP_VERY_LOW_VALUE",
	2: "SMP_LOW_VALUE",
	3: "SMP_NORMAL_VALUE",
	4: "SMP_HIGH_VALUE",
	5: "SMP_VERY_HIGH_VALUE",
}

var SgMessagePriority_value = map[string]int32{
	"SMP_UNKNOWN":         0,
	"SMP_VERY_LOW_VALUE":  1,
	"SMP_LOW_VALUE":       2,
	"SMP_NORMAL_VALUE":    3,
	"SMP_HIGH_VALUE":      4,
	"SMP_VERY_HIGH_VALUE": 5,
}

func (x SgMessagePriority) String() string {
	return proto.EnumName(SgMessagePriority_name, int32(x))
}

func (SgMessagePriority) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_94955c5bcc4b00ce, []int{2}
}

type SgMessageReq struct {
	// 消息字节数组
	Message *SgMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty" validate:"required"`
	// 用户ID
	UserIdList []int64 `protobuf:"varint,2,rep,packed,name=user_id_list,json=userIdList,proto3" json:"user_id_list,omitempty"`
	// 文件的ID
	FileId               string   `protobuf:"bytes,3,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SgMessageReq) Reset()         { *m = SgMessageReq{} }
func (m *SgMessageReq) String() string { return proto.CompactTextString(m) }
func (*SgMessageReq) ProtoMessage()    {}
func (*SgMessageReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_94955c5bcc4b00ce, []int{0}
}
func (m *SgMessageReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SgMessageReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SgMessageReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SgMessageReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SgMessageReq.Merge(m, src)
}
func (m *SgMessageReq) XXX_Size() int {
	return m.Size()
}
func (m *SgMessageReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SgMessageReq.DiscardUnknown(m)
}

var xxx_messageInfo_SgMessageReq proto.InternalMessageInfo

type SgMessageResp struct {
	// 状态
	Status int64 `protobuf:"varint,1,opt,name=status,proto3" json:"status"`
	// 发送的终端数
	Clients              int64    `protobuf:"varint,2,opt,name=clients,proto3" json:"clients"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SgMessageResp) Reset()         { *m = SgMessageResp{} }
func (m *SgMessageResp) String() string { return proto.CompactTextString(m) }
func (*SgMessageResp) ProtoMessage()    {}
func (*SgMessageResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_94955c5bcc4b00ce, []int{1}
}
func (m *SgMessageResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SgMessageResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SgMessageResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SgMessageResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SgMessageResp.Merge(m, src)
}
func (m *SgMessageResp) XXX_Size() int {
	return m.Size()
}
func (m *SgMessageResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SgMessageResp.DiscardUnknown(m)
}

var xxx_messageInfo_SgMessageResp proto.InternalMessageInfo

// 消息定义
type SgMessage struct {
	// 消息类别
	Dmt SgMessageType `protobuf:"varint,1,opt,name=dmt,proto3,enum=work.sg.common.v1.SgMessageType" json:"dmt,omitempty"`
	// 自定义的消息体
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	// 创建时间
	CreateAt int64 `protobuf:"varint,3,opt,name=create_at,json=createAt,proto3" json:"create_at,omitempty"`
	// 过期时间
	ExpireAt int64 `protobuf:"varint,4,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty"`
	// 需要发送的时间
	SendAt int64 `protobuf:"varint,5,opt,name=send_at,json=sendAt,proto3" json:"send_at,omitempty"`
	// 通知登记
	Priority SgMessagePriority `protobuf:"varint,6,opt,name=priority,proto3,enum=work.sg.common.v1.SgMessagePriority" json:"priority,omitempty"`
	// 事件类别
	Event                SgEvent  `protobuf:"varint,7,opt,name=event,proto3,enum=work.sg.common.v1.SgEvent" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SgMessage) Reset()         { *m = SgMessage{} }
func (m *SgMessage) String() string { return proto.CompactTextString(m) }
func (*SgMessage) ProtoMessage()    {}
func (*SgMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_94955c5bcc4b00ce, []int{2}
}
func (m *SgMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SgMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SgMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SgMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SgMessage.Merge(m, src)
}
func (m *SgMessage) XXX_Size() int {
	return m.Size()
}
func (m *SgMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_SgMessage.DiscardUnknown(m)
}

var xxx_messageInfo_SgMessage proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("work.sg.common.v1.SgMessageType", SgMessageType_name, SgMessageType_value)
	proto.RegisterEnum("work.sg.common.v1.SgEvent", SgEvent_name, SgEvent_value)
	proto.RegisterEnum("work.sg.common.v1.SgMessagePriority", SgMessagePriority_name, SgMessagePriority_value)
	proto.RegisterType((*SgMessageReq)(nil), "work.sg.common.v1.SgMessageReq")
	proto.RegisterType((*SgMessageResp)(nil), "work.sg.common.v1.SgMessageResp")
	proto.RegisterType((*SgMessage)(nil), "work.sg.common.v1.SgMessage")
}

func init() { proto.RegisterFile("work/sg/common/v1/api.proto", fileDescriptor_94955c5bcc4b00ce) }

var fileDescriptor_94955c5bcc4b00ce = []byte{
	// 674 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcf, 0x4e, 0xdb, 0x4a,
	0x14, 0xc6, 0x71, 0x9c, 0x3f, 0xe4, 0x04, 0xb8, 0x66, 0x40, 0x97, 0x08, 0x50, 0x88, 0xa2, 0x7b,
	0x25, 0x84, 0x44, 0x7c, 0x81, 0xdd, 0x95, 0x2a, 0xd5, 0x49, 0xa6, 0x24, 0xaa, 0xf3, 0x47, 0xe3,
	0x00, 0x2a, 0x9b, 0x91, 0xc1, 0x53, 0xd7, 0x6a, 0x12, 0x07, 0xcf, 0x24, 0x2d, 0xaf, 0xd1, 0x25,
	0x8f, 0xd1, 0xa7, 0x60, 0xd9, 0x27, 0x40, 0x2d, 0x4b, 0x96, 0x7d, 0x82, 0x6a, 0xc6, 0x4e, 0x1a,
	0x89, 0x8a, 0x95, 0xe7, 0xfc, 0xbe, 0xef, 0x9c, 0xf9, 0x66, 0x2c, 0x1b, 0x76, 0x3e, 0x85, 0xd1,
	0x47, 0x93, 0xfb, 0xe6, 0x75, 0x38, 0x1c, 0x86, 0x23, 0x73, 0x7a, 0x64, 0xba, 0xe3, 0xa0, 0x3a,
	0x8e, 0x42, 0x11, 0xa2, 0x75, 0x29, 0x56, 0xb9, 0x5f, 0x8d, 0xc5, 0xea, 0xf4, 0x68, 0x7b, 0xd3,
	0x0f, 0xfd, 0x50, 0xa9, 0xa6, 0x5c, 0xc5, 0xc6, 0xed, 0x5d, 0x3f, 0x0c, 0xfd, 0x01, 0x93, 0xad,
	0xa6, 0x3b, 0x1a, 0x85, 0xc2, 0x15, 0x41, 0x38, 0xe2, 0xb1, 0x5a, 0xb9, 0xd3, 0x60, 0xc5, 0xf1,
	0xdb, 0x8c, 0x73, 0xd7, 0x67, 0x84, 0xdd, 0xa0, 0x2e, 0xe4, 0x86, 0x71, 0x55, 0xd4, 0xca, 0xda,
	0x7e, 0xe1, 0x78, 0xb7, 0xfa, 0x6c, 0xa7, 0xea, 0xbc, 0xa3, 0xb6, 0xf5, 0xf3, 0x61, 0x6f, 0x63,
	0xea, 0x0e, 0x02, 0xcf, 0x15, 0xec, 0xff, 0x4a, 0xc4, 0x6e, 0x26, 0x41, 0xc4, 0xbc, 0x0a, 0x99,
	0x4d, 0x41, 0x65, 0x58, 0x99, 0x70, 0x16, 0xd1, 0xc0, 0xa3, 0x83, 0x80, 0x8b, 0x62, 0xaa, 0xac,
	0xef, 0xeb, 0x04, 0x24, 0x6b, 0x79, 0x76, 0xc0, 0x05, 0xda, 0x82, 0xdc, 0xfb, 0x60, 0xc0, 0x68,
	0xe0, 0x15, 0xf5, 0xb2, 0xb6, 0x9f, 0x27, 0x59, 0x59, 0xb6, 0xbc, 0xca, 0x25, 0xac, 0x2e, 0x64,
	0xe3, 0x63, 0x54, 0x81, 0x2c, 0x17, 0xae, 0x98, 0x70, 0x95, 0x4d, 0xaf, 0xc1, 0xd3, 0xc3, 0x5e,
	0x42, 0x48, 0xf2, 0x44, 0xff, 0x42, 0xee, 0x7a, 0x10, 0xb0, 0x91, 0xe0, 0xc5, 0x94, 0x32, 0x15,
	0x9e, 0x1e, 0xf6, 0x66, 0x88, 0xcc, 0x16, 0x95, 0xbb, 0x14, 0xe4, 0xe7, 0xc3, 0xd1, 0x31, 0xe8,
	0xde, 0x50, 0xa8, 0xa9, 0x6b, 0xc7, 0xe5, 0x97, 0x4e, 0xdc, 0xbf, 0x1d, 0x33, 0x22, 0xcd, 0x08,
	0x41, 0xda, 0x73, 0x85, 0xab, 0x76, 0xc9, 0x13, 0xb5, 0x46, 0x3b, 0x90, 0xbf, 0x8e, 0x98, 0x2b,
	0x18, 0x75, 0x85, 0x3a, 0x8c, 0x4e, 0x96, 0x63, 0x60, 0x09, 0x29, 0xb2, 0xcf, 0xe3, 0x20, 0x52,
	0x62, 0x3a, 0x16, 0x63, 0x60, 0xa9, 0x4b, 0xe0, 0x6c, 0xe4, 0x49, 0x29, 0xa3, 0xa4, 0xac, 0x2c,
	0x2d, 0x81, 0x5e, 0xc3, 0xf2, 0x38, 0x0a, 0xc2, 0x28, 0x10, 0xb7, 0xc5, 0xac, 0xca, 0xf7, 0xcf,
	0x4b, 0xf9, 0x7a, 0x89, 0x97, 0xcc, 0xbb, 0xd0, 0x7f, 0x90, 0x61, 0x53, 0x36, 0x12, 0xc5, 0x9c,
	0x6a, 0xdf, 0xfe, 0x63, 0x3b, 0x96, 0x0e, 0x12, 0x1b, 0x0f, 0x5e, 0x2d, 0x5c, 0xbc, 0x3c, 0x30,
	0xfa, 0x0b, 0x0a, 0x8d, 0x76, 0x9f, 0x9e, 0x75, 0xde, 0x76, 0xba, 0x17, 0x1d, 0x63, 0x09, 0x15,
	0x20, 0x27, 0x41, 0xbf, 0xd6, 0x30, 0x34, 0xb4, 0x06, 0xa0, 0x0a, 0x82, 0x3b, 0x0d, 0xc7, 0x48,
	0x1d, 0x7c, 0xd5, 0x20, 0x97, 0x4c, 0x94, 0x9a, 0x83, 0x17, 0x1a, 0x11, 0xac, 0x39, 0x98, 0x5a,
	0x8d, 0x06, 0xb5, 0x1c, 0x5a, 0xef, 0xda, 0xb6, 0xa1, 0x25, 0xac, 0x81, 0xed, 0x39, 0x4b, 0x25,
	0x7d, 0xf5, 0xe6, 0x29, 0xed, 0xe1, 0xb6, 0xa1, 0x23, 0x80, 0xac, 0x83, 0x69, 0xdf, 0x79, 0x63,
	0xa4, 0x67, 0x33, 0x7a, 0xbd, 0xb9, 0x3f, 0x23, 0x13, 0x3a, 0x98, 0x76, 0xfa, 0xb4, 0xde, 0xc6,
	0x9d, 0xbe, 0x91, 0x45, 0x06, 0xac, 0xc4, 0x80, 0xe0, 0xf3, 0x96, 0x83, 0x8d, 0x1c, 0xda, 0x85,
	0x62, 0x62, 0xe9, 0xda, 0x0e, 0xa6, 0xbd, 0xb3, 0x9a, 0xdd, 0xaa, 0x53, 0xa7, 0x69, 0x11, 0x6c,
	0x2c, 0x1f, 0x7c, 0xd1, 0x60, 0xfd, 0xd9, 0x2d, 0xaa, 0xb1, 0xed, 0xde, 0x42, 0xfe, 0xbf, 0x01,
	0x49, 0x70, 0x8e, 0xc9, 0x3b, 0x6a, 0x77, 0x2f, 0xe8, 0xb9, 0x65, 0x9f, 0x61, 0x43, 0x43, 0xeb,
	0xb0, 0x2a, 0xf9, 0x6f, 0x94, 0x42, 0x9b, 0x60, 0x48, 0xd4, 0xe9, 0x92, 0xb6, 0x65, 0x27, 0x54,
	0x57, 0xe1, 0xdb, 0x3d, 0xda, 0x6c, 0x9d, 0x36, 0x13, 0x96, 0x46, 0x5b, 0xb0, 0x31, 0x1f, 0xba,
	0x20, 0x64, 0x6a, 0xf8, 0xfe, 0x47, 0x69, 0xe9, 0xfe, 0xb1, 0xa4, 0x7d, 0x7b, 0x2c, 0x69, 0xdf,
	0x1f, 0x4b, 0xda, 0xe5, 0x89, 0x1f, 0x88, 0x0f, 0x93, 0x2b, 0xf9, 0xee, 0x4c, 0x2e, 0xc2, 0x68,
	0x6a, 0x7a, 0xfc, 0xea, 0x90, 0xb3, 0x68, 0xca, 0xa2, 0xc3, 0xf8, 0xe3, 0x7f, 0xf6, 0xdb, 0xb8,
	0xca, 0x2a, 0xe1, 0xe4, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0xb1, 0x15, 0x67, 0x52, 0x04,
	0x00, 0x00,
}

func (m *SgMessageReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SgMessageReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Message != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Message.Size()))
		n1, err := m.Message.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.UserIdList) > 0 {
		dAtA3 := make([]byte, len(m.UserIdList)*10)
		var j2 int
		for _, num1 := range m.UserIdList {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	if len(m.FileId) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.FileId)))
		i += copy(dAtA[i:], m.FileId)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SgMessageResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SgMessageResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Status != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Status))
	}
	if m.Clients != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Clients))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *SgMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SgMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Dmt != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Dmt))
	}
	if len(m.Data) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.CreateAt != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.CreateAt))
	}
	if m.ExpireAt != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.ExpireAt))
	}
	if m.SendAt != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.SendAt))
	}
	if m.Priority != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Priority))
	}
	if m.Event != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Event))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SgMessageReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Message != nil {
		l = m.Message.Size()
		n += 1 + l + sovApi(uint64(l))
	}
	if len(m.UserIdList) > 0 {
		l = 0
		for _, e := range m.UserIdList {
			l += sovApi(uint64(e))
		}
		n += 1 + sovApi(uint64(l)) + l
	}
	l = len(m.FileId)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SgMessageResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Status != 0 {
		n += 1 + sovApi(uint64(m.Status))
	}
	if m.Clients != 0 {
		n += 1 + sovApi(uint64(m.Clients))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *SgMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Dmt != 0 {
		n += 1 + sovApi(uint64(m.Dmt))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovApi(uint64(l))
	}
	if m.CreateAt != 0 {
		n += 1 + sovApi(uint64(m.CreateAt))
	}
	if m.ExpireAt != 0 {
		n += 1 + sovApi(uint64(m.ExpireAt))
	}
	if m.SendAt != 0 {
		n += 1 + sovApi(uint64(m.SendAt))
	}
	if m.Priority != 0 {
		n += 1 + sovApi(uint64(m.Priority))
	}
	if m.Event != 0 {
		n += 1 + sovApi(uint64(m.Event))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SgMessageReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SgMessageReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SgMessageReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Message == nil {
				m.Message = &SgMessage{}
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowApi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.UserIdList = append(m.UserIdList, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowApi
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthApi
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthApi
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.UserIdList) == 0 {
					m.UserIdList = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowApi
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.UserIdList = append(m.UserIdList, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field UserIdList", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SgMessageResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SgMessageResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SgMessageResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clients", wireType)
			}
			m.Clients = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Clients |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SgMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SgMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SgMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Dmt", wireType)
			}
			m.Dmt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Dmt |= SgMessageType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApi
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateAt", wireType)
			}
			m.CreateAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreateAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpireAt", wireType)
			}
			m.ExpireAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpireAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendAt", wireType)
			}
			m.SendAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SendAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Priority", wireType)
			}
			m.Priority = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Priority |= SgMessagePriority(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			m.Event = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Event |= SgEvent(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthApi
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)
