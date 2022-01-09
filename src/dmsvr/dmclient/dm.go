// Code generated by goctl. DO NOT EDIT!
// Source: dm.proto

package dmclient

import (
	"context"

	"github.com/go-things/things/src/dmsvr/dm"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccessAuthReq            = dm.AccessAuthReq
	DeviceData               = dm.DeviceData
	DeviceInfo               = dm.DeviceInfo
	GetDeviceDataReq         = dm.GetDeviceDataReq
	GetDeviceDataResp        = dm.GetDeviceDataResp
	GetDeviceInfoReq         = dm.GetDeviceInfoReq
	GetDeviceInfoResp        = dm.GetDeviceInfoResp
	GetDeviceLogReq          = dm.GetDeviceLogReq
	GetDeviceLogResp         = dm.GetDeviceLogResp
	GetProductInfoReq        = dm.GetProductInfoReq
	GetProductInfoResp       = dm.GetProductInfoResp
	GetProductTemplateReq    = dm.GetProductTemplateReq
	LoginAuthReq             = dm.LoginAuthReq
	ManResp                  = dm.ManResp
	ManageDeviceReq          = dm.ManageDeviceReq
	ManageProductReq         = dm.ManageProductReq
	ManageProductTemplateReq = dm.ManageProductTemplateReq
	PageInfo                 = dm.PageInfo
	ProductInfo              = dm.ProductInfo
	ProductTemplate          = dm.ProductTemplate
	Response                 = dm.Response
	SendActionReq            = dm.SendActionReq
	SendActionResp           = dm.SendActionResp
	SendPropertyReq          = dm.SendPropertyReq
	SendPropertyResp         = dm.SendPropertyResp

	Dm interface {
		// 登录认证
		LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error)
		// 操作认证
		AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error)
		// 设备管理
		ManageDevice(ctx context.Context, in *ManageDeviceReq, opts ...grpc.CallOption) (*DeviceInfo, error)
		// 产品管理
		ManageProduct(ctx context.Context, in *ManageProductReq, opts ...grpc.CallOption) (*ProductInfo, error)
		// 获取产品信息
		GetProductInfo(ctx context.Context, in *GetProductInfoReq, opts ...grpc.CallOption) (*GetProductInfoResp, error)
		// 产品模板管理
		ManageProductTemplate(ctx context.Context, in *ManageProductTemplateReq, opts ...grpc.CallOption) (*ProductTemplate, error)
		// 获取产品信息
		GetProductTemplate(ctx context.Context, in *GetProductTemplateReq, opts ...grpc.CallOption) (*ProductTemplate, error)
		// 获取设备信息
		GetDeviceInfo(ctx context.Context, in *GetDeviceInfoReq, opts ...grpc.CallOption) (*GetDeviceInfoResp, error)
		// 获取设备最新数据
		GetDeviceData(ctx context.Context, in *GetDeviceDataReq, opts ...grpc.CallOption) (*GetDeviceDataResp, error)
		// 获取设备日志信息
		GetDeviceLog(ctx context.Context, in *GetDeviceLogReq, opts ...grpc.CallOption) (*GetDeviceLogResp, error)
		// 同步调用设备行为
		SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error)
		// 同步调用设备属性
		SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error)
	}

	defaultDm struct {
		cli zrpc.Client
	}
)

func NewDm(cli zrpc.Client) Dm {
	return &defaultDm{
		cli: cli,
	}
}

// 登录认证
func (m *defaultDm) LoginAuth(ctx context.Context, in *LoginAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.LoginAuth(ctx, in, opts...)
}

// 操作认证
func (m *defaultDm) AccessAuth(ctx context.Context, in *AccessAuthReq, opts ...grpc.CallOption) (*Response, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.AccessAuth(ctx, in, opts...)
}

// 设备管理
func (m *defaultDm) ManageDevice(ctx context.Context, in *ManageDeviceReq, opts ...grpc.CallOption) (*DeviceInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ManageDevice(ctx, in, opts...)
}

// 产品管理
func (m *defaultDm) ManageProduct(ctx context.Context, in *ManageProductReq, opts ...grpc.CallOption) (*ProductInfo, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ManageProduct(ctx, in, opts...)
}

// 获取产品信息
func (m *defaultDm) GetProductInfo(ctx context.Context, in *GetProductInfoReq, opts ...grpc.CallOption) (*GetProductInfoResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetProductInfo(ctx, in, opts...)
}

// 产品模板管理
func (m *defaultDm) ManageProductTemplate(ctx context.Context, in *ManageProductTemplateReq, opts ...grpc.CallOption) (*ProductTemplate, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.ManageProductTemplate(ctx, in, opts...)
}

// 获取产品信息
func (m *defaultDm) GetProductTemplate(ctx context.Context, in *GetProductTemplateReq, opts ...grpc.CallOption) (*ProductTemplate, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetProductTemplate(ctx, in, opts...)
}

// 获取设备信息
func (m *defaultDm) GetDeviceInfo(ctx context.Context, in *GetDeviceInfoReq, opts ...grpc.CallOption) (*GetDeviceInfoResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetDeviceInfo(ctx, in, opts...)
}

// 获取设备最新数据
func (m *defaultDm) GetDeviceData(ctx context.Context, in *GetDeviceDataReq, opts ...grpc.CallOption) (*GetDeviceDataResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetDeviceData(ctx, in, opts...)
}

// 获取设备日志信息
func (m *defaultDm) GetDeviceLog(ctx context.Context, in *GetDeviceLogReq, opts ...grpc.CallOption) (*GetDeviceLogResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.GetDeviceLog(ctx, in, opts...)
}

// 同步调用设备行为
func (m *defaultDm) SendAction(ctx context.Context, in *SendActionReq, opts ...grpc.CallOption) (*SendActionResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.SendAction(ctx, in, opts...)
}

// 同步调用设备属性
func (m *defaultDm) SendProperty(ctx context.Context, in *SendPropertyReq, opts ...grpc.CallOption) (*SendPropertyResp, error) {
	client := dm.NewDmClient(m.cli.Conn())
	return client.SendProperty(ctx, in, opts...)
}
