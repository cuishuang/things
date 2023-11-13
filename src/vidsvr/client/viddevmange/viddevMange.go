// Code generated by goctl. DO NOT EDIT.
// Source: vid.proto

package viddevmange

import (
	"context"

	"github.com/i-Things/things/src/vidsvr/internal/svc"
	"github.com/i-Things/things/src/vidsvr/pb/vid"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	PageInfo            = vid.PageInfo
	PageInfo_OrderBy    = vid.PageInfo_OrderBy
	Point               = vid.Point
	Response            = vid.Response
	VidPlanInfo         = vid.VidPlanInfo
	ViddevInfo          = vid.ViddevInfo
	VidmgrInfo          = vid.VidmgrInfo
	VidmgrInfoCountReq  = vid.VidmgrInfoCountReq
	VidmgrInfoCountResp = vid.VidmgrInfoCountResp
	VidmgrInfoDeleteReq = vid.VidmgrInfoDeleteReq
	VidmgrInfoIndexReq  = vid.VidmgrInfoIndexReq
	VidmgrInfoIndexResp = vid.VidmgrInfoIndexResp
	VidmgrInfoReadReq   = vid.VidmgrInfoReadReq
	VidrecordfileInfo   = vid.VidrecordfileInfo
	VidstreamInfo       = vid.VidstreamInfo

	ViddevMange interface {
		// 流添加
		VideoaddInfroCreate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
		// 流更新
		VideoaddInfroUpdate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
		// 删除流
		VideoaddInfroDelete(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
		// 获取流列表
		VideoaddInfroIndex(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
		// 获取流信息详情
		VideoaddInfroRead(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
		// 统计流 在线，离线，未激活
		VideoaddInfroCount(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error)
	}

	defaultViddevMange struct {
		cli zrpc.Client
	}

	directViddevMange struct {
		svcCtx *svc.ServiceContext
		svr    vid.ViddevMangeServer
	}
)

func NewViddevMange(cli zrpc.Client) ViddevMange {
	return &defaultViddevMange{
		cli: cli,
	}
}

func NewDirectViddevMange(svcCtx *svc.ServiceContext, svr vid.ViddevMangeServer) ViddevMange {
	return &directViddevMange{
		svr:    svr,
		svcCtx: svcCtx,
	}
}

// 流添加
func (m *defaultViddevMange) VideoaddInfroCreate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewViddevMangeClient(m.cli.Conn())
	return client.VideoaddInfroCreate(ctx, in, opts...)
}

// 流添加
func (d *directViddevMange) VideoaddInfroCreate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VideoaddInfroCreate(ctx, in)
}

// 流更新
func (m *defaultViddevMange) VideoaddInfroUpdate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewViddevMangeClient(m.cli.Conn())
	return client.VideoaddInfroUpdate(ctx, in, opts...)
}

// 流更新
func (d *directViddevMange) VideoaddInfroUpdate(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VideoaddInfroUpdate(ctx, in)
}

// 删除流
func (m *defaultViddevMange) VideoaddInfroDelete(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewViddevMangeClient(m.cli.Conn())
	return client.VideoaddInfroDelete(ctx, in, opts...)
}

// 删除流
func (d *directViddevMange) VideoaddInfroDelete(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VideoaddInfroDelete(ctx, in)
}

// 获取流列表
func (m *defaultViddevMange) VideoaddInfroIndex(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewViddevMangeClient(m.cli.Conn())
	return client.VideoaddInfroIndex(ctx, in, opts...)
}

// 获取流列表
func (d *directViddevMange) VideoaddInfroIndex(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VideoaddInfroIndex(ctx, in)
}

// 获取流信息详情
func (m *defaultViddevMange) VideoaddInfroRead(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewViddevMangeClient(m.cli.Conn())
	return client.VideoaddInfroRead(ctx, in, opts...)
}

// 获取流信息详情
func (d *directViddevMange) VideoaddInfroRead(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VideoaddInfroRead(ctx, in)
}

// 统计流 在线，离线，未激活
func (m *defaultViddevMange) VideoaddInfroCount(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	client := vid.NewViddevMangeClient(m.cli.Conn())
	return client.VideoaddInfroCount(ctx, in, opts...)
}

// 统计流 在线，离线，未激活
func (d *directViddevMange) VideoaddInfroCount(ctx context.Context, in *ViddevInfo, opts ...grpc.CallOption) (*Response, error) {
	return d.svr.VideoaddInfroCount(ctx, in)
}
