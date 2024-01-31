// Code generated by goctl. DO NOT EDIT.
// Source: vid.proto

package vidmgrstreammanage

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

	VidmgrStreamManage interface {
		// 流添加
		VidmgrStreamCreate(ctx context.Context, in *VidmgrStreamCreateReq, opts ...grpc.CallOption) (*Response, error)
		// 流更新
		VidmgrStreamUpdate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error)
		// 删除流
		VidmgrStreamDelete(ctx context.Context, in *VidmgrStreamDeleteReq, opts ...grpc.CallOption) (*Response, error)
		// 获取流列表
		VidmgrStreamIndex(ctx context.Context, in *VidmgrStreamIndexReq, opts ...grpc.CallOption) (*VidmgrStreamIndexResp, error)
		// 获取流信息详情
		VidmgrStreamRead(ctx context.Context, in *VidmgrStreamReadReq, opts ...grpc.CallOption) (*VidmgrStream, error)
		// 统计流 在线，离线，未激活
		VidmgrStreamCount(ctx context.Context, in *VidmgrStreamCountReq, opts ...grpc.CallOption) (*VidmgrStreamCountResp, error)
	}

	defaultVidmgrStreamManage struct {
		cli zrpc.Client
	}

	directVidmgrStreamManage struct {
		svcCtx *svc.ServiceContext
		svr    vid.VidmgrStreamManageServer
	}
)

func NewVidmgrStreamManage(cli zrpc.Client) VidmgrStreamManage {
	return &defaultVidmgrStreamManage{
		cli: cli,
	}
}

func NewDirectVidmgrStreamManage(svcCtx *svc.ServiceContext, svr vid.VidmgrStreamManageServer) VidmgrStreamManage {
	return &directVidmgrStreamManage{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 流添加
func (m *defaultVidmgrStreamManage) VidmgrStreamCreate(ctx context.Context, in *VidmgrStreamCreateReq, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrStreamManageClient(m.cli.Conn())
	return client.VidmgrStreamCreate(ctx, in, opts...)
}

// 流添加
func (d *directVidmgrStreamManage) VidmgrStreamCreate(ctx context.Context, in *VidmgrStreamCreateReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrStreamCreate(ctx, in)
}

// 流更新
func (m *defaultVidmgrStreamManage) VidmgrStreamUpdate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrStreamManageClient(m.cli.Conn())
	return client.VidmgrStreamUpdate(ctx, in, opts...)
}

// 流更新
func (d *directVidmgrStreamManage) VidmgrStreamUpdate(ctx context.Context, in *VidmgrStream, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrStreamUpdate(ctx, in)
}

// 删除流
func (m *defaultVidmgrStreamManage) VidmgrStreamDelete(ctx context.Context, in *VidmgrStreamDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewVidmgrStreamManageClient(m.cli.Conn())
	return client.VidmgrStreamDelete(ctx, in, opts...)
}

// 删除流
func (d *directVidmgrStreamManage) VidmgrStreamDelete(ctx context.Context, in *VidmgrStreamDeleteReq, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VidmgrStreamDelete(ctx, in)
}

// 获取流列表
func (m *defaultVidmgrStreamManage) VidmgrStreamIndex(ctx context.Context, in *VidmgrStreamIndexReq, opts ...grpc.CallOption) (*VidmgrStreamIndexResp, error) {
	client := vid.NewVidmgrStreamManageClient(m.cli.Conn())
	return client.VidmgrStreamIndex(ctx, in, opts...)
}

// 获取流列表
func (d *directVidmgrStreamManage) VidmgrStreamIndex(ctx context.Context, in *VidmgrStreamIndexReq, opts ...grpc.CallOption) (*VidmgrStreamIndexResp, error) {
	return d.svr.VidmgrStreamIndex(ctx, in)
}

// 获取流信息详情
func (m *defaultVidmgrStreamManage) VidmgrStreamRead(ctx context.Context, in *VidmgrStreamReadReq, opts ...grpc.CallOption) (*VidmgrStream, error) {
	client := vid.NewVidmgrStreamManageClient(m.cli.Conn())
	return client.VidmgrStreamRead(ctx, in, opts...)
}

// 获取流信息详情
func (d *directVidmgrStreamManage) VidmgrStreamRead(ctx context.Context, in *VidmgrStreamReadReq, opts ...grpc.CallOption) (*VidmgrStream, error) {
	return d.svr.VidmgrStreamRead(ctx, in)
}

// 统计流 在线，离线，未激活
func (m *defaultVidmgrStreamManage) VidmgrStreamCount(ctx context.Context, in *VidmgrStreamCountReq, opts ...grpc.CallOption) (*VidmgrStreamCountResp, error) {
	client := vid.NewVidmgrStreamManageClient(m.cli.Conn())
	return client.VidmgrStreamCount(ctx, in, opts...)
}

// 统计流 在线，离线，未激活
func (d *directVidmgrStreamManage) VidmgrStreamCount(ctx context.Context, in *VidmgrStreamCountReq, opts ...grpc.CallOption) (*VidmgrStreamCountResp, error) {
	return d.svr.VidmgrStreamCount(ctx, in)
}
