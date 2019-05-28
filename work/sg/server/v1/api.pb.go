// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: work/sg/server/v1/api.proto

// package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

package v1

import (
	context "context"
	fmt "fmt"
	_ "github.com/bilibili/kratos/tool/protobuf/pkg/extensions/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/storv/dsb-server-proto/work/sg/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

func init() { proto.RegisterFile("work/sg/server/v1/api.proto", fileDescriptor_df3e0e30dc35b6a3) }

var fileDescriptor_df3e0e30dc35b6a3 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0xcf, 0x2f, 0xca,
	0xd6, 0x2f, 0x4e, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2f, 0x33, 0xd4, 0x4f, 0x2c,
	0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4c, 0xce, 0xcf, 0xd5, 0x2b, 0x4e, 0xd7,
	0x03, 0xc9, 0x65, 0x26, 0xa7, 0xea, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7, 0xa7, 0xe7, 0x83, 0x65,
	0xf5, 0x41, 0x2c, 0x88, 0x42, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x90, 0x56, 0xfd,
	0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xa8, 0x2c, 0xdc, 0x8e, 0xe4,
	0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0x14, 0x3b, 0x8c, 0xea, 0xb8, 0xd8, 0x5c, 0x8a, 0x93, 0x1c, 0x0b,
	0x32, 0x85, 0x4a, 0xb8, 0xb8, 0x83, 0x53, 0xf3, 0x52, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53,
	0x85, 0xe4, 0xf5, 0x40, 0xda, 0x40, 0xd6, 0x43, 0xb4, 0xe9, 0x95, 0x19, 0xea, 0x05, 0xa7, 0x43,
	0x65, 0x83, 0x52, 0x0b, 0xa5, 0x14, 0xf0, 0x2b, 0x28, 0x2e, 0x50, 0x52, 0x6a, 0xba, 0xfc, 0x64,
	0x32, 0x93, 0x8c, 0x92, 0x94, 0x7e, 0x4a, 0x71, 0x92, 0x2e, 0xc2, 0x87, 0xc5, 0x08, 0x6b, 0x9c,
	0x5c, 0x4f, 0x3c, 0x94, 0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0xa3, 0x8c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0x40, 0xa6, 0xea, 0x17, 0x97, 0xe4,
	0x17, 0x95, 0x21, 0xe9, 0xd6, 0x85, 0x78, 0x1d, 0x23, 0xd0, 0x92, 0xd8, 0xc0, 0x12, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xb1, 0x4e, 0x98, 0xb0, 0x50, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DsbApiClient is the client API for DsbApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DsbApiClient interface {
	SendMessage(ctx context.Context, in *v1.SgMessageReq, opts ...grpc.CallOption) (*v1.SgMessageResp, error)
}

type dsbApiClient struct {
	cc *grpc.ClientConn
}

func NewDsbApiClient(cc *grpc.ClientConn) DsbApiClient {
	return &dsbApiClient{cc}
}

func (c *dsbApiClient) SendMessage(ctx context.Context, in *v1.SgMessageReq, opts ...grpc.CallOption) (*v1.SgMessageResp, error) {
	out := new(v1.SgMessageResp)
	err := c.cc.Invoke(ctx, "/com.sg.service.v1.DsbApi/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DsbApiServer is the server API for DsbApi service.
type DsbApiServer interface {
	SendMessage(context.Context, *v1.SgMessageReq) (*v1.SgMessageResp, error)
}

func RegisterDsbApiServer(s *grpc.Server, srv DsbApiServer) {
	s.RegisterService(&_DsbApi_serviceDesc, srv)
}

func _DsbApi_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.SgMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DsbApiServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.sg.service.v1.DsbApi/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DsbApiServer).SendMessage(ctx, req.(*v1.SgMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _DsbApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.sg.service.v1.DsbApi",
	HandlerType: (*DsbApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _DsbApi_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "work/sg/server/v1/api.proto",
}
