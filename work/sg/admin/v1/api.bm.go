// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: work/sg/admin/v1/api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	work/sg/admin/v1/api.proto
*/
package v1

import (
	"context"
	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
	google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
	work_sg_common_v1 "github.com/storv/dsb-server-proto/work/sg/common/v1"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDsbAdminApiSendMessage = "/dsb-admin/v1/sendMessage"
var PathDsbAdminApiStatClient = "/dsb-admin/v1/stat"

// DsbAdminApiBMServer is the server API for DsbAdminApi service.
type DsbAdminApiBMServer interface {
	SendMessage(ctx context.Context, req *work_sg_common_v1.SgMessageReq) (resp *work_sg_common_v1.SgMessageResp, err error)

	StatClient(ctx context.Context, req *google_protobuf1.Empty) (resp *work_sg_common_v1.CountClientResp, err error)
}

var v1DsbAdminApiSvc DsbAdminApiBMServer

func dsbAdminApiSendMessage(c *bm.Context) {
	p := new(work_sg_common_v1.SgMessageReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DsbAdminApiSvc.SendMessage(c, p)
	c.JSON(resp, err)
}

func dsbAdminApiStatClient(c *bm.Context) {
	p := new(google_protobuf1.Empty)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1DsbAdminApiSvc.StatClient(c, p)
	c.JSON(resp, err)
}

// RegisterDsbAdminApiBMServer Register the blademaster route
func RegisterDsbAdminApiBMServer(e *bm.Engine, server DsbAdminApiBMServer) {
	v1DsbAdminApiSvc = server
	e.POST("/dsb-admin/v1/sendMessage", dsbAdminApiSendMessage)
	e.POST("/dsb-admin/v1/stat", dsbAdminApiStatClient)
}
