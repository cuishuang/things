// Code generated by goctl. DO NOT EDIT.
// Source: vid.proto

package vidmgrconfigmanage

import (
	"context"

	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/pb/vid"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PageInfo              = vid.PageInfo
	PageInfo_OrderBy      = vid.PageInfo_OrderBy
	Response              = vid.Response
	StreamTrack           = vid.StreamTrack
	VidPlanInfo           = vid.VidPlanInfo
	VidmgrConfig          = vid.VidmgrConfig
	VidmgrConfigDeleteReq = vid.VidmgrConfigDeleteReq
	VidmgrConfigIndexReq  = vid.VidmgrConfigIndexReq
	VidmgrConfigIndexResp = vid.VidmgrConfigIndexResp
	VidmgrConfigReadReq   = vid.VidmgrConfigReadReq
	VidmgrInfo            = vid.VidmgrInfo
	VidmgrInfoActiveReq   = vid.VidmgrInfoActiveReq
	VidmgrInfoCountReq    = vid.VidmgrInfoCountReq
	VidmgrInfoCountResp   = vid.VidmgrInfoCountResp
	VidmgrInfoDeleteReq   = vid.VidmgrInfoDeleteReq
	VidmgrInfoIndexReq    = vid.VidmgrInfoIndexReq
	VidmgrInfoIndexResp   = vid.VidmgrInfoIndexResp
	VidmgrInfoReadReq     = vid.VidmgrInfoReadReq
	VidmgrStream          = vid.VidmgrStream
	VidmgrStreamCountReq  = vid.VidmgrStreamCountReq
	VidmgrStreamCountResp = vid.VidmgrStreamCountResp
	VidmgrStreamCreateReq = vid.VidmgrStreamCreateReq
	VidmgrStreamDeleteReq = vid.VidmgrStreamDeleteReq
	VidmgrStreamIndexReq  = vid.VidmgrStreamIndexReq
	VidmgrStreamIndexResp = vid.VidmgrStreamIndexResp
	VidmgrStreamReadReq   = vid.VidmgrStreamReadReq
	VidrecordfileInfo     = vid.VidrecordfileInfo

	VidmgrConfigManage interface {
		// 新建配置
		VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error)
		// 删除配置
		VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 更新配置
		VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error)
		// 配置列表
		VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigIndexReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error)
		// 获取配置信息详情
		VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error)
	}

	defaultVidmgrConfigManage struct {
		cli zrpc.Client
	}

	directVidmgrConfigManage struct {
		svcCtx *svc.ServiceContext
		svr    vid.VidmgrConfigManageServer
	}
)

func NewVidmgrConfigManage(cli zrpc.Client) VidmgrConfigManage {
	return &defaultVidmgrConfigManage{
		cli: cli,
	}
}

func NewDirectVidmgrConfigManage(svcCtx *svc.ServiceContext, svr vid.VidmgrConfigManageServer) VidmgrConfigManage {
	return &directVidmgrConfigManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 新建配置
func (m *defaultVidmgrConfigManage) VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrConfigManageClient(m.cli.Conn())
	return client.VidmgrConfigCreate(ctx, in, opts...)
}

// 新建配置
func (d *directVidmgrConfigManage) VidmgrConfigCreate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrConfigCreate(ctx, in)
}

// 删除配置
func (m *defaultVidmgrConfigManage) VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrConfigManageClient(m.cli.Conn())
	return client.VidmgrConfigDelete(ctx, in, opts...)
}

// 删除配置
func (d *directVidmgrConfigManage) VidmgrConfigDelete(ctx context.Context, in *VidmgrConfigDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrConfigDelete(ctx, in)
}

// 更新配置
func (m *defaultVidmgrConfigManage) VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrConfigManageClient(m.cli.Conn())
	return client.VidmgrConfigUpdate(ctx, in, opts...)
}

// 更新配置
func (d *directVidmgrConfigManage) VidmgrConfigUpdate(ctx context.Context, in *VidmgrConfig, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrConfigUpdate(ctx, in)
}

// 配置列表
func (m *defaultVidmgrConfigManage) VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigIndexReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error) {
	client := vid.NewVidmgrConfigManageClient(m.cli.Conn())
	return client.VidmgrConfigIndex(ctx, in, opts...)
}

// 配置列表
func (d *directVidmgrConfigManage) VidmgrConfigIndex(ctx context.Context, in *VidmgrConfigIndexReq, opts ...grpc.CallOption) (*VidmgrConfigIndexResp, error) {
	return d.svr.VidmgrConfigIndex(ctx, in)
}

// 获取配置信息详情
func (m *defaultVidmgrConfigManage) VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error) {
	client := vid.NewVidmgrConfigManageClient(m.cli.Conn())
	return client.VidmgrConfigRead(ctx, in, opts...)
}

// 获取配置信息详情
func (d *directVidmgrConfigManage) VidmgrConfigRead(ctx context.Context, in *VidmgrConfigReadReq, opts ...grpc.CallOption) (*VidmgrConfig, error) {
	return d.svr.VidmgrConfigRead(ctx, in)
}
