// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type DmServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedDmServer
}

func NewDmServer(svcCtx *svc.ServiceContext) *DmServer {
	return &DmServer{
		svcCtx: svcCtx,
	}
}

// 设备登录认证
func (s *DmServer) LoginAuth(ctx context.Context, in *dm.LoginAuthReq) (*dm.Response, error) {
	l := logic.NewLoginAuthLogic(ctx, s.svcCtx)
	return l.LoginAuth(in)
}

// 设备操作认证
func (s *DmServer) AccessAuth(ctx context.Context, in *dm.AccessAuthReq) (*dm.Response, error) {
	l := logic.NewAccessAuthLogic(ctx, s.svcCtx)
	return l.AccessAuth(in)
}

// 鉴定是否是root账号
func (s *DmServer) RootCheck(ctx context.Context, in *dm.RootCheckReq) (*dm.Response, error) {
	l := logic.NewRootCheckLogic(ctx, s.svcCtx)
	return l.RootCheck(in)
}

// 设备管理
func (s *DmServer) ManageDevice(ctx context.Context, in *dm.ManageDeviceReq) (*dm.DeviceInfo, error) {
	l := logic.NewManageDeviceLogic(ctx, s.svcCtx)
	return l.ManageDevice(in)
}

// 产品管理
func (s *DmServer) ManageProduct(ctx context.Context, in *dm.ManageProductReq) (*dm.ProductInfo, error) {
	l := logic.NewManageProductLogic(ctx, s.svcCtx)
	return l.ManageProduct(in)
}

// 获取产品信息
func (s *DmServer) GetProductInfo(ctx context.Context, in *dm.GetProductInfoReq) (*dm.GetProductInfoResp, error) {
	l := logic.NewGetProductInfoLogic(ctx, s.svcCtx)
	return l.GetProductInfo(in)
}

// 产品模板管理
func (s *DmServer) ManageProductTemplate(ctx context.Context, in *dm.ManageProductTemplateReq) (*dm.ProductTemplate, error) {
	l := logic.NewManageProductTemplateLogic(ctx, s.svcCtx)
	return l.ManageProductTemplate(in)
}

// 获取产品信息
func (s *DmServer) GetProductTemplate(ctx context.Context, in *dm.GetProductTemplateReq) (*dm.ProductTemplate, error) {
	l := logic.NewGetProductTemplateLogic(ctx, s.svcCtx)
	return l.GetProductTemplate(in)
}

// 获取设备信息
func (s *DmServer) GetDeviceInfo(ctx context.Context, in *dm.GetDeviceInfoReq) (*dm.GetDeviceInfoResp, error) {
	l := logic.NewGetDeviceInfoLogic(ctx, s.svcCtx)
	return l.GetDeviceInfo(in)
}

// 获取设备调试信息记录登入登出,操作
func (s *DmServer) GetDeviceDescribeLog(ctx context.Context, in *dm.GetDeviceDescribeLogReq) (*dm.GetDeviceDescribeLogResp, error) {
	l := logic.NewGetDeviceDescribeLogLogic(ctx, s.svcCtx)
	return l.GetDeviceDescribeLog(in)
}

// 获取设备数据信息
func (s *DmServer) GetDeviceData(ctx context.Context, in *dm.GetDeviceDataReq) (*dm.GetDeviceDataResp, error) {
	l := logic.NewGetDeviceDataLogic(ctx, s.svcCtx)
	return l.GetDeviceData(in)
}

// 同步调用设备行为
func (s *DmServer) SendAction(ctx context.Context, in *dm.SendActionReq) (*dm.SendActionResp, error) {
	l := logic.NewSendActionLogic(ctx, s.svcCtx)
	return l.SendAction(in)
}

// 同步调用设备属性
func (s *DmServer) SendProperty(ctx context.Context, in *dm.SendPropertyReq) (*dm.SendPropertyResp, error) {
	l := logic.NewSendPropertyLogic(ctx, s.svcCtx)
	return l.SendProperty(in)
}
