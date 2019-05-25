// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: work/sg/server/v1/api.proto

/*
Package work_sg_server is a generated blademaster stub package.
This code was generated with kratos/tool/protobuf/protoc-gen-bm v0.1.

package 命名使用 {appid}.{version} 的方式, version 形如 v1, v2 ..

It is generated from these files:
	work/sg/server/v1/api.proto
*/
package work_sg_server

import (
	"context"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
	"github.com/bilibili/kratos/pkg/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathDsbApiSendMessage = "/dsbServer/sendMessage"

// DsbApiBMServer is the server API for DsbApi service.
type DsbApiBMServer interface {
	SendMessage(ctx context.Context, req *DsbMessageReq) (resp *DsbMessageResp, err error)
}

var DsbApiSvc DsbApiBMServer

func dsbApiSendMessage(c *bm.Context) {
	p := new(DsbMessageReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := DsbApiSvc.SendMessage(c, p)
	c.JSON(resp, err)
}

// RegisterDsbApiBMServer Register the blademaster route
func RegisterDsbApiBMServer(e *bm.Engine, server DsbApiBMServer) {
	DsbApiSvc = server
	e.POST("/dsbServer/sendMessage", dsbApiSendMessage)
}
